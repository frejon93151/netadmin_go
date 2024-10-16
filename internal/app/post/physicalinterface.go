/*
Copyright © 2024 HALMSTADS STADSNÄT AB fredrik.jonsson1@halmstad.se
*/
package post

import (
	"fmt"
	"net/http"
	"os"

	"github.com/frejon93151/netadmin_go/internal/app/utils"
)

func PhysIfConnect(id int, addressId int) (resp *http.Response, err error) {
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"https://login.halmstadsstadsnat.se/api/%s/%d/connect?addressId=%d",
			"physicalinterfaces",
			id,
			addressId,
		),
		nil,
	)
	if err != nil {
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("NETADMIN__ACCESS_TOKEN"))
	req.Header.Add("Content-Type", "application/json")

	resp, err = http.DefaultClient.Do(req)
	if resp.StatusCode == 401 {
		utils.RenewAccessToken()
		resp, err = PhysIfConnect(id, addressId)
	}
	return
}

func PhysIfDisconnect(id int) (resp *http.Response, err error) {
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"https://login.halmstadsstadsnat.se/api/%s/%d/disconnect",
			"physicalinterfaces",
			id,
		),
		nil,
	)

	if err != nil {
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("NETADMIN__ACCESS_TOKEN"))
	req.Header.Add("Content-Type", "application/json")

	resp, err = http.DefaultClient.Do(req)
	if resp.StatusCode == 401 {
		utils.RenewAccessToken()
		resp, err = PhysIfDisconnect(id)
	}
	return
}
