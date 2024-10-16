/*
Copyright © 2024 HALMSTADS STADSNÄT AB fredrik.jonsson1@halmstad.se
*/
package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"time"
	"unicode"
)

const (
	charComment       = '#'
	prefixSingleQuote = '\''
	prefixDoubleQuote = '"'
	exportPrefix      = "export"
)

func ParseEnv(buf bytes.Buffer) (out map[string]string, err error) {
	out = make(map[string]string)
	err = parseEnvBytes(buf.Bytes(), out)
	return
}

func parseEnvBytes(src []byte, out map[string]string) error {
	src = bytes.Replace(src, []byte("\r\n"), []byte("\n"), -1)
	cutset := src
	for {
		cutset = getStatementStart(cutset)
		if cutset == nil {
			break
		}

		key, left, err := locateKeyName(cutset)
		if err != nil {
			return err
		}

		value, left, err := extractVarValue(left, out)
		if err != nil {
			return err
		}
		out[key] = value
		cutset = left
	}
	return nil
}

func locateKeyName(src []byte) (key string, cutset []byte, err error) {
	src = bytes.TrimLeftFunc(src, isSpace)
	if bytes.HasPrefix(src, []byte(exportPrefix)) {
		trimmed := bytes.TrimPrefix(src, []byte(exportPrefix))
		if bytes.IndexFunc(trimmed, isSpace) == 0 {
			src = bytes.TrimLeftFunc(trimmed, isSpace)
		}
	}

	offset := 0

loop:
	for i, char := range src {
		rchar := rune(char)
		if isSpace(rchar) {
			continue
		}

		switch char {
		case '=', ':':
			key = string(src[0:i])
			offset = i + 1
			break loop
		case '_':
		default:
			if unicode.IsLetter(rchar) || unicode.IsNumber(rchar) || rchar == '.' {
				continue
			}

			return "", nil, fmt.Errorf(
				`unexpected character %q in variable name near %q`,
				string(char), string(src),
			)
		}
	}

	if len(src) == 0 {
		return "", nil, errors.New("zero length string")
	}

	key = strings.TrimRightFunc(key, unicode.IsSpace)
	cutset = bytes.TrimLeftFunc(src[offset:], isSpace)
	return key, cutset, nil
}

func extractVarValue(src []byte, vars map[string]string) (value string, rest []byte, err error) {
	quote, hasPrefix := hasQuotePrefix(src)
	if !hasPrefix {
		endOfLine := bytes.IndexFunc(src, isLineEnd)

		if endOfLine == -1 {
			endOfLine = len(src)
			if endOfLine == 0 {
				return "", nil, nil
			}
		}

		line := []rune(string(src[0:endOfLine]))

		endOfVar := len(line)
		if endOfVar == 0 {
			return "", src[endOfLine:], nil
		}

		for i := endOfVar - 1; i >= 0; i-- {
			if line[i] == charComment && i > 0 {
				if isSpace(line[i-1]) {
					endOfVar = i
					break
				}
			}
		}
		trimmed := strings.TrimFunc(string(line[0:endOfVar]), isSpace)

		return expandVariables(trimmed, vars), src[endOfLine:], nil
	}

	for i := 1; i < len(src); i++ {
		if char := src[i]; char != quote {
			continue
		}
		if prevChar := src[i-1]; prevChar == '\\' {
			continue
		}

		trimFunc := isCharFunc(rune(quote))
		value = string(bytes.TrimLeftFunc(bytes.TrimRightFunc(src[0:i], trimFunc), trimFunc))
		if quote == prefixDoubleQuote {
			value = expandVariables(expandEscapes(value), vars)
		}
		return value, src[i+1:], nil
	}

	valEndIndex := bytes.IndexFunc(src, isCharFunc('\n'))
	if valEndIndex == -1 {
		valEndIndex = len(src)
	}

	return "", nil, fmt.Errorf("unterminated quoted value %s", src[:valEndIndex])
}

func expandEscapes(str string) string {
	out := escapeRegex.ReplaceAllStringFunc(str, func(match string) string {
		c := strings.TrimPrefix(match, `\`)
		switch c {
		case "n":
			return "\n"
		case "r":
			return "\r"
		default:
			return match
		}
	})
	return unescapeCharsRegex.ReplaceAllString(out, "$1")
}

func hasQuotePrefix(src []byte) (prefix byte, isQuored bool) {
	if len(src) == 0 {
		return 0, false
	}

	switch prefix := src[0]; prefix {
	case prefixDoubleQuote, prefixSingleQuote:
		return prefix, true
	default:
		return 0, false
	}
}

func isCharFunc(char rune) func(rune) bool {
	return func(v rune) bool {
		return v == char
	}
}

func isSpace(r rune) bool {
	switch r {
	case '\t', '\v', '\f', ' ', 0x85, 0xA0:
		return true
	}
	return false
}

func isLineEnd(r rune) bool {
	if r == '\n' || r == '\r' {
		return true
	}
	return false
}

var (
	escapeRegex        = regexp.MustCompile(`\\.`)
	expandVarRegex     = regexp.MustCompile(`(\\)?(\$)(\()?\{?([A-Z0-9_]+)?\}?`)
	unescapeCharsRegex = regexp.MustCompile(`\\([^$])`)
)

func expandVariables(v string, m map[string]string) string {
	return expandVarRegex.ReplaceAllStringFunc(v, func(s string) string {
		submatch := expandVarRegex.FindStringSubmatch(s)

		if submatch == nil {
			return s
		}
		if submatch[1] == "\\" || submatch[2] == "(" {
			return submatch[0][1:]
		} else if submatch[4] != "" {
			return m[submatch[4]]
		}
		return s
	})
}

// getStatementPosition returns position of statement begin.
//
// It skips any comment line or non-whitespace character.
func getStatementStart(src []byte) []byte {
	pos := indexOfNonSpaceChar(src)
	if pos == -1 {
		return nil
	}

	src = src[pos:]
	if src[0] != charComment {
		return src
	}

	// skip comment section
	pos = bytes.IndexFunc(src, isCharFunc('\n'))
	if pos == -1 {
		return nil
	}

	return getStatementStart(src[pos:])
}

func indexOfNonSpaceChar(src []byte) int {
	return bytes.IndexFunc(src, func(r rune) bool {
		return !unicode.IsSpace(r)
	})
}

func ReadFile(filename string) (buf bytes.Buffer, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = io.Copy(&buf, file)
	return
}

func Check(i interface{}, e error) {
	if e != nil {
		fmt.Printf("%s error: %s\n", NameOf(i), e)
		os.Exit(1)
	}
}

func NameOf(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func RenewAccessToken() {
	data := url.Values{}
	data.Add("client_id", os.Getenv("NETADMIN__CLIENT_ID"))
	data.Add("client_secret", os.Getenv("NETADMIN__CLIENT_SECRET"))
	data.Add("grant_type", os.Getenv("NETADMIN__GRANT_TYPE"))

	res, err := http.PostForm(
		"https://login.halmstadsstadsnat.se/oauth2/token",
		data,
	)
	if err != nil {
		log.Fatalf(
			"%s: Failed getting accesstoken - %s",
			time.Now().Format("RFC1123"),
			err.Error())
		return
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf(
			"%s: Failed to read http.Response.Body - %s",
			time.Now().Format("RFC1123"),
			err.Error())
		return
	} else if res.StatusCode > 299 {
		log.Fatalf(
			"%s: %d %s - %s",
			time.Now().Format("RFC1123"),
			res.StatusCode,
			res.Status,
			string(body))
		return
	}
	content := make(map[string]any)
	json.Unmarshal(body, &content)
	accesstoken := fmt.Sprint(content["access_token"])
	os.Setenv("NETADMIN__ACCESS_TOKEN", accesstoken)
}

func ExclusiveParams(a any, b any) bool {
	if a != nil && b != nil {
		return false
	}
	return true
}

func MapStruct[T interface{}](obj interface{}) (newObj T, err error) {
	objT := reflect.TypeOf(obj)
	var buf []byte
	if objT.ConvertibleTo(reflect.TypeOf(newObj)) {
		buf, err = json.Marshal(obj)
		if err != nil {
			return
		}
		err = json.Unmarshal(buf, &newObj)
		if err != nil {
			return
		}
	}
	return newObj, nil
}

func TryAddParams(params *url.Values, key string, value any) {
	kind := reflect.TypeOf(value).Kind()
	switch kind {
	case reflect.String:
		if value != "" && value != "0" && value != "false" {
			(*params).Add(key, fmt.Sprintf("%s", value))
		}
	case reflect.Int:
		if value != 0 {
			(*params).Add(key, fmt.Sprintf("%d", value))
		}
	case reflect.Bool:
		if value != false {
			(*params).Add(key, fmt.Sprintf("%t", value))
		}
	}
}
