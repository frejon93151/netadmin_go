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
var urlTempl = baseUrl + "%s?%s"
var phys = "physicalinterfaces"
var devices = "devices"
var devicePhysTempl = "devices/%d/physicalinterfaces"

func doGet(opts doGetOpts) (*http.Response, error) {
	req := &http.Request{}
	var err error
	if opts.req != nil {
		req = opts.req
		req.Method = "GET"
	} else {
		req, err = http.NewRequest(
			"GET",
			endpointBuilder(opts.endpoint, opts.params),
			nil,
		)
		if err != nil {
			return nil, err
		}
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", os.Getenv("NETADMIN__ACCESS_TOKEN"))

	res, err := http.DefaultClient.Do(req)
	resp := *res
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 401 {
		utils.RenewAccessToken()
		doGet(opts)
	}

	return &resp, err
}

func endpointBuilder(endpoint string, params *url.Values) string {
	completeUrl := fmt.Sprintf(urlTempl, endpoint, params.Encode())
	return completeUrl
}

type doGetOpts struct {
	endpoint string
	params   *url.Values
	req      *http.Request
}
