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

var (
	devices            = "devices"
	devicePhysIfTempl  = "devices/%d/physicalinterfaces"
	devicePhysLnkTempl = "devices/%d/physicallinks"
	deviceAuthTempl    = "devices/%d/authentications"
)

func Devices(opts models.DeviceGetOpts) (resp *http.Response, err error) {
	params, err := deviceParams(opts)
	if err != nil {
		return nil, err
	}

	resp, err = doGet(devices, &params)
	return
}

func DevicePhysicalInterfaces(id int, pageIndex int, itemsPerPage int) (resp *http.Response, err error) {
	params := &url.Values{
		"pageIndex":    {fmt.Sprintf("%d", pageIndex)},
		"itemsPerPage": {fmt.Sprintf("%d", itemsPerPage)},
	}

	resp, err = doGet(fmt.Sprintf(devicePhysIfTempl, id), params)
	return
}

func DevicePhysicalLinks(id int, pageIndex int, itemsPerPage int) (resp *http.Response, err error) {
	params := &url.Values{
		"pageIndex":    {fmt.Sprintf("%d", pageIndex)},
		"itemsPerPage": {fmt.Sprintf("%d", itemsPerPage)},
	}

	resp, err = doGet(fmt.Sprintf(devicePhysLnkTempl, id), params)
	return
}

func DeviceAuthentications(id int) (resp *http.Response, err error) {
	resp, err = doGet(fmt.Sprintf(deviceAuthTempl, id), nil)
	return
}

func deviceParams(opts models.DeviceGetOpts) (params url.Values, err error) {
	params = make(url.Values)
	if exclusive := utils.ExclusiveParams(opts.ManagementAddress, opts.ManagementAddresses); exclusive {
		err = fmt.Errorf(
			"deviceParams: Error - %v and %v are mutually exclusive options",
			utils.NameOf(opts.ManagementAddress),
			utils.NameOf(opts.ManagementAddresses))
		return
	}
	for _, item := range opts.Ids {
		utils.TryAddParams(&params, "ids", item)
	}
	for _, item := range opts.CompanyIds {
		utils.TryAddParams(&params, "companyIds", fmt.Sprintf("%d", item))
	}
	for _, item := range opts.ManagementAddresses {
		utils.TryAddParams(&params, "managementAddresses", item)
	}
	utils.TryAddParams(&params, "name", opts.Name)
	utils.TryAddParams(&params, "managementAddress", opts.ManagementAddress)
	utils.TryAddParams(&params, "hostName", opts.HostName)
	utils.TryAddParams(&params, "serialNumber", opts.SerialNumber)
	utils.TryAddParams(&params, "function", opts.DeviceFunction)
	utils.TryAddParams(&params, "type", opts.TypeName)
	utils.TryAddParams(&params, "definitionName", opts.DefinitionName)
	utils.TryAddParams(&params, "software", opts.Software)
	utils.TryAddParams(&params, "alertGroup", opts.AlertGroup)
	utils.TryAddParams(&params, "administrativeStatus", opts.AdministrativeStatus)
	utils.TryAddParams(&params, "operationalStatus", opts.OperationalStatus)
	utils.TryAddParams(&params, "modifiedSince", opts.ModifiedSince)
	utils.TryAddParams(&params, "pageIndex", fmt.Sprintf("%d", opts.PageIndex))
	utils.TryAddParams(&params, "itemsPerPage", fmt.Sprintf("%d", opts.ItemsPerPage))
	utils.TryAddParams(&params, "onlyRootDevices", fmt.Sprintf("%t", opts.OnlyRootDevices))
	return
}
