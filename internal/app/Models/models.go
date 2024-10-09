/*
Copyright © 2024 HALMSTADS STADSNÄT AB fredrik.jonsson1@halmstad.se
*/
package Models

/*
Paginated response
*/
type FilterResp[T comparable] struct {
	PageIndex        int `json:"pageIndex"`
	ItemsPerPage     int `json:"itemsPerPage"`
	TotalItems       int `json:"totalItems"`
	StartIndex       int `json:"startIndex"`
	CurrentItemCount int `json:"currentItemCount"`
	Items            T   `json:"items"`
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
	IncludeCredentials       *bool  `json:"includeCredentials,omitempty"`
	Name                     string `json:"name,omitempty"`
	ManagementAddress        string `json:"managementAddress,omitempty"`
	HostName                 string `json:"hostName,omitempty"`
	ReferenceName            string `json:"referenceName,omitempty"`
	SerialNumber             string `json:"serialNumber,omitempty"`
	AdministrativeStatus     string `json:"administrativeStatus,omitempty"`
	OperationalStatus        string `json:"operationalStatus,omitempty"`
	CustomerId               int    `json:"customerId,omitempty"`
	SiteId                   int    `json:"siteId,omitempty"`
	AddressId                int    `json:"addressId,omitempty"`
	DeviceDefinitionsVersion string `json:"deviceDefinitionsVersion,omitempty"`
	Software                 string `json:"software,omitempty"`
	SoftwareVersion          string `json:"softwareVersion,omitempty"`
}

type DeviceGetOpts struct {
	Ids                  []int    `json:"ids,omitempty"`
	CompanyIds           []int    `json:"companyIds,omitempty"`
	ManagementAddresses  []string `json:"managementAddresses,omitempty"`
	Name                 string   `json:"name,omitempty"`
	ManagementAddress    string   `json:"managementAddress,omitempty"`
	HostName             string   `json:"hostName,omitempty"`
	SerialNumber         string   `json:"serialNumber,omitempty"`
	DeviceFunction       string   `json:"deviceFunction,omitempty"`
	TypeName             string   `json:"typeName,omitempty"`
	DefinitionName       string   `json:"definitionName,omitempty"`
	Software             string   `json:"software,omitempty"`
	AlertGroup           string   `json:"alertGroup,omitempty"`
	AdministrativeStatus string   `json:"administrativeStatus,omitempty"`
	OperationalStatus    string   `json:"operationalStatus,omitempty"`
	ModifiedSince        string   `json:"modifiedSince,omitempty"`
	PageIndex            int      `json:"pageIndex,omitempty"`
	ItemsPerPage         int      `json:"itemsPerPage,omitempty"`
	OnlyRootDevices      bool     `json:"onlyRootDevices,omitempty"`
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
	Ids          []int  `json:"ids,omitempty"`
	Name         string `json:"name,omitempty"`
	DeviceName   string `json:"deviceName,omitempty"`
	Identifier   string `json:"identifier,omitempty"`
	Option82     string `json:"option82,omitempty"`
	Fqpn         string `json:"fqpn,omitempty"`
	Number       int    `json:"number,omitempty"`
	DeviceId     int    `json:"deviceId,omitempty"`
	AddressId    int    `json:"addressId,omitempty"`
	PageIndex    int    `json:"pageIndex,omitempty"`
	ItemsPerPage int    `json:"itemsPerPage,omitempty"`
}
