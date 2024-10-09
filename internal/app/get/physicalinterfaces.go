/*
Copyright © 2024 HALMSTADS STADSNÄT AB fredrik.jonsson1@halmstad.se
*/
package get

import (
	"fmt"
	"net/http"
	"net/url"

	models "github.com/frejon93151/netadmin_go/internal/app/Models"
)

func PhysicalInterfaces(opts models.PhysGetOpts) (resp *http.Response, err error) {
	params, err := physParams(opts)
	if err != nil {
		return
	}

	resp, err = doGet(doGetOpts{endpoint: phys, params: params})
	return
}

func physParams(opts models.PhysGetOpts) (params *url.Values, err error) {
	for _, item := range opts.Ids {
		params.Add("ids", fmt.Sprintf("%d", item))
	}
	params.Add("name", opts.Name)
	params.Add("deviceName", opts.DeviceName)
	params.Add("identifier", opts.Identifier)
	params.Add("option82", opts.Option82)
	params.Add("fqpn", opts.Fqpn)
	params.Add("number", fmt.Sprintf("%d", opts.Number))
	params.Add("deviceId", fmt.Sprintf("%d", opts.DeviceId))
	params.Add("addressId", fmt.Sprintf("%d", opts.AddressId))
	params.Add("pageIndex", fmt.Sprintf("%d", opts.PageIndex))
	params.Add("itemsPerPage", fmt.Sprintf("%d", opts.ItemsPerPage))
	return
}
