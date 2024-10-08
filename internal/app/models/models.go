/*
Copyright © 2024 HALMSTADS STADSNÄT AB fredrik.jonsson1@halmstad.se
*/
package models

/*
Paginated response
*/
type FilterResp[T comparable] struct {
	PageIndex        int `json:"pageIndex"`
	ItemsPerPage     int `json:"itemsPerPage"`
	TotalItems       int `json:"totalItems"`
	StartIndex       int `json:"startIndex"`
	CurrentItemCount int `json:"currentItemCount"`
}

/*
Device structs
*/
type DeviceDTO struct {
	Id                   int    `json:"id"`
	CompanyId            int    `json:"companyId"`
	Name                 string `json:"name"`
	Reference            string `json:"reference"`
	ManagementAddress    string `json:"managementAddress"`
	HostName             string `json:"hostName"`
	SerialNumber         string `json:"serialNumber"`
	Function             string `json:"function"`
	DeviceDefinitionId   int    `json:"deviceDefinitionId"`
	DeviceDefinition     string `json:"deviceDefinition"`
	Software             string `json:"software"`
	SoftwareVersion      string `json:"softwareVersion"`
	Site                 string `json:"site"`
	AlertGroup           string `json:"alertGroup"`
	Customer             string `json:"customer"`
	Description          string `json:"description"`
	AdministrativeStatus string `json:"administrativeStatus"`
	OperationalStatus    string `json:"operationalStatus"`
	Changed              string `json:"changed"`
	Created              string `json:"created"`
	CustomerId           int    `json:"customerId"`
	SiteId               int    `json:"siteId"`
	AddressId            int    `json:"addressId"`
}

type DevicePostOpts struct {
	IncludeCredentials       *bool   `json:"includeCredentials"`
	Name                     *string `json:"name"`
	ManagementAddress        *string `json:"managementAddress"`
	HostName                 *string `json:"hostName"`
	ReferenceName            *string `json:"referenceName"`
	SerialNumber             *string `json:"serialNumber"`
	AdministrativeStatus     *string `json:"administrativeStatus"`
	OperationalStatus        *string `json:"operationalStatus"`
	CustomerId               *int    `json:"customerId"`
	SiteId                   *int    `json:"siteId"`
	AddressId                *int    `json:"addressId"`
	DeviceDefinitionsVersion *string `json:"deviceDefinitionsVersion"`
	Software                 *string `json:"software"`
	SoftwareVersion          *string `json:"softwareVersion"`
}

type DeviceGetOpts struct {
	Ids                  []int    `json:"ids"`
	CompanyIds           []int    `json:"companyIds"`
	ManagementAddresses  []string `json:"managementAddresses"`
	Name                 string   `json:"name"`
	ManagementAddress    string   `json:"managementAddress"`
	HostName             string   `json:"hostName"`
	SerialNumber         string   `json:"serialNumber"`
	DeviceFunction       string   `json:"deviceFunction"`
	TypeName             string   `json:"typeName"`
	DefinitionName       string   `json:"definitionName"`
	Software             string   `json:"software"`
	AlertGroup           string   `json:"alertGroup"`
	AdministrativeStatus string   `json:"administrativeStatus"`
	OperationalStatus    string   `json:"operationalStatus"`
	ModifiedSince        string   `json:"modifiedSince"`
	PageIndex            int      `json:"pageIndex"`
	ItemsPerPage         int      `json:"itemsPerPage"`
	OnlyRootDevices      bool     `json:"onlyRootDevices"`
}

/*
PhysicalInterface structs
*/
type PhysDTO struct {
	Id                   int    `json:"id"`
	Name                 string `json:"name"`
	Number               int    `json:"number"`
	Panel                string `json:"panel"`
	Medium               string `json:"medium"`
	Role                 string `json:"role"`
	Identifier           string `json:"identifier"`
	Description          string `json:"description"`
	Option82             string `json:"option82"`
	Fqpn                 string `json:"fqpn"`
	AddressId            int    `json:"addressId"`
	AdministrativeStatus string `json:"administrativeStatus"`
	OperationalStatus    string `json:"operationalStatus"`
	SlotName             string `json:"slotName"`
	DeviceId             int    `json:"deviceId"`
	DeviceName           string `json:"deviceName"`
	RootDeviceId         int    `json:"rootDeviceId"`
	RootDeviceName       string `json:"rootDeviceName"`
}

type PhysGetOpts struct {
	Ids          []int  `json:"ids"`
	Name         string `json:"name"`
	DeviceName   string `json:"deviceName"`
	Identifier   string `json:"identifier"`
	Option82     string `json:"option82"`
	Fqpn         string `json:"fqpn"`
	Number       int    `json:"number"`
	DeviceId     int    `json:"deviceId"`
	AddressId    int    `json:"addressId"`
	PageIndex    int    `json:"pageIndex"`
	ItemsPerPage int    `json:"itemsPerPage"`
}
