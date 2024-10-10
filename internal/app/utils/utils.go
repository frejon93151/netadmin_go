/*
Copyright © 2024 HALMSTADS STADSNÄT AB fredrik.jonsson1@halmstad.se
*/
package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"runtime"
	"time"
)

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
	var content map[string]interface{}
	json.Unmarshal(body, &content)
	os.Setenv("NETADMIN__ACCESS_TOKEN", fmt.Sprint(content["access_token"]))
}

func ExclusiveParams(a any, b any) bool {
	if a != nil && b != nil {
		return true
	}
	return false
}

func MapStruct[T interface{}](obj interface{}) (newObj T) {
	objT := reflect.TypeOf(obj)
	fmt.Printf("%v", objT.ConvertibleTo(reflect.TypeOf(newObj)))
	return
}
