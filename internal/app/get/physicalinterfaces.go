/*
Copyright © 2024 HALMSTADS STADSNÄT AB fredrik.jonsson1@halmstad.se
*/
package get

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/frejon93151/netadmin_go/internal/app/models"
	"github.com/frejon93151/netadmin_go/internal/app/utils"
)

var phys = "physicalinterfaces"

func PhysicalInterfaces(opts models.PhysGetOpts) (*http.Response, error) {
	params, err := physParams(opts)
	if err != nil {
		return nil, err
	}

	res, err := doGet(phys, &params)
	resp := *res
	return &resp, err
}

func physParams(opts models.PhysGetOpts) (params url.Values, err error) {
	params = make(url.Values)
	for _, item := range opts.Ids {
		utils.TryAddParams(&params, "ids", fmt.Sprintf("%d", item))
	}
	utils.TryAddParams(&params, "name", opts.Name)
	utils.TryAddParams(&params, "deviceName", opts.DeviceName)
	utils.TryAddParams(&params, "identifier", opts.Identifier)
	utils.TryAddParams(&params, "option82", opts.Option82)
	utils.TryAddParams(&params, "fqpn", opts.Fqpn)
	utils.TryAddParams(&params, "number", fmt.Sprintf("%d", opts.Number))
	utils.TryAddParams(&params, "deviceId", fmt.Sprintf("%d", opts.DeviceId))
	utils.TryAddParams(&params, "addressId", fmt.Sprintf("%d", opts.AddressId))
	utils.TryAddParams(&params, "pageIndex", fmt.Sprintf("%d", opts.PageIndex))
	utils.TryAddParams(&params, "itemsPerPage", fmt.Sprintf("%d", opts.ItemsPerPage))
	return
}
