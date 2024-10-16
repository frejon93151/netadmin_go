/*
Copyright © 2024 HALMSTADS STADSNÄT AB fredrik.jonsson1@halmstad.se
*/
package get

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/frejon93151/netadmin_go/internal/app/utils"
)

var baseUrl = "https://login.halmstadsstadsnat.se/api/"
var urlTempl = baseUrl + "%s"

var retries = 0

func doGet(endpoint string, params *url.Values) (*http.Response, error) {
	accesstoken, hasToken := os.LookupEnv("NETADMIN__ACCESS_TOKEN")
	if !hasToken || accesstoken != "" {
		utils.RenewAccessToken()
		accesstoken, hasToken = os.LookupEnv("NETADMIN__ACCESS_TOKEN")
	}

	uri := endpointBuilder(endpoint, params)
	req, err := http.NewRequest(
		"GET",
		uri,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accesstoken))

	fmt.Println(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 401 && retries < 5 {
		retries++
		utils.RenewAccessToken()
		resp, err = doGet(endpoint, params)
	}

	retries = 0

	return resp, err
}

func endpointBuilder(endpoint string, params *url.Values) string {
	uri := fmt.Sprintf(urlTempl, endpoint)
	if params != nil {
		if len(params.Encode()) != 0 {
			uri = uri + "?" + params.Encode()
		}
	}
	return uri
}

type doGetOpts struct {
	endpoint string
	params   *url.Values
}
