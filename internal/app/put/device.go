/*
Copyright © 2024 HALMSTADS STADSNÄT AB fredrik.jonsson1@halmstad.se
*/
package put

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/frejon93151/netadmin_go/internal/app/models"
	"github.com/frejon93151/netadmin_go/internal/app/utils"
)

func Device(id int, opts models.DevicePostOpts) (*http.Response, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf("https://login.halmstadsstadsnat.se/api/%s/%d", "devices", id),
		bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("NETADMIN__ACCESS_TOKEN")))
	req.Header.Add("Content-Type", "application/json")

	fmt.Sprintln(string(body))
	fmt.Sprintln(req)

	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode == 401 {
		utils.RenewAccessToken()
		resp, err = Device(id, opts)
	}
	return resp, err
}
