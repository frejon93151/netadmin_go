/*
Copyright © 2024 HALMSTADS STADSNÄT AB fredrik.jonsson1@halmstad.se
*/
package get

import (
	"fmt"
	"net/http"
	"net/url"

	models "github.com/frejon93151/netadmin_go/internal/app/models"
	"github.com/frejon93151/netadmin_go/internal/app/utils"
)

func Devices(opts models.DeviceGetOpts) (resp *http.Response, err error) {
	params, err := deviceParams(opts)
	if err != nil {
		return
	}

	resp, err = doGet(doGetOpts{
		endpoint: devices,
		params:   params,
	})
	return
}

func DevicePhysicalInterfaces(id int, pageIndex *int, itemsPerPage *int) (resp *http.Response, err error) {
	params := &url.Values{
		"pageIndex":    {fmt.Sprintf("%d", pageIndex)},
		"itemsPerPage": {fmt.Sprintf("%d", itemsPerPage)},
	}

	resp, err = doGet(doGetOpts{
		endpoint: fmt.Sprintf(devicePhysTempl, id),
		params:   params,
	})
	return
}

func deviceParams(opts models.DeviceGetOpts) (params *url.Values, err error) {
	if utils.ExclusiveParams(opts.ManagementAddress, opts.ManagementAddresses) {
		err = fmt.Errorf(
			"%s and %s are mutually exclusive options",
			utils.NameOf(opts.ManagementAddress),
			utils.NameOf(opts.ManagementAddresses))
		return
	}
	for _, item := range *opts.Ids {
		params.Add("ids", fmt.Sprintf("%d", item))
	}
	for _, item := range *opts.CompanyIds {
		params.Add("companyIds", fmt.Sprintf("%d", item))
	}
	for _, item := range *opts.ManagementAddresses {
		params.Add("managementAddresses", item)
	}
	params.Add("name", opts.Name)
	params.Add("managementAddress", opts.ManagementAddress)
	params.Add("hostName", opts.HostName)
	params.Add("serialNumber", opts.SerialNumber)
	params.Add("function", opts.DeviceFunction)
	params.Add("type", opts.TypeName)
	params.Add("definitionName", opts.DefinitionName)
	params.Add("software", opts.Software)
	params.Add("alertGroup", opts.AlertGroup)
	params.Add("administrativeStatus", opts.AdministrativeStatus)
	params.Add("operationalStatus", opts.OperationalStatus)
	params.Add("modifiedSince", opts.ModifiedSince)
	params.Add("pageIndex", fmt.Sprintf("%d", opts.PageIndex))
	params.Add("itemsPerPage", fmt.Sprintf("%d", opts.ItemsPerPage))
	params.Add("onlyRootDevices", fmt.Sprintf("%t", opts.OnlyRootDevices))
	return
}
