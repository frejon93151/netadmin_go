/*
Copyright © 2024 HALMSTADS STADSNÄT AB fredrik.jonsson1@halmstad.se
*/
package models

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
type PhysPostOpts struct {
	Role                 string `json:"role,omitempty"`
	Panel                string `json:"panel,omitempty"`
	Option82             string `json:"option82,omitempty"`
	OperationalStatus    string `json:"operationalStatus,omitempty"`
	AdministrativeStatus string `json:"administrativeStatus,omitempty"`
	Identifier           string `json:"identifier,omitempty"`
	Fqpn                 string `json:"fqpn,omitempty"`
	Description          string `json:"description,omitempty"`
}
