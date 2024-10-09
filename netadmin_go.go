/*
Copyright © 2024 HALMSTADS STADSNÄT AB fredrik.jonsson1@halmstad.se
*/
package netadmin_go

import (
	"net/http"

	"github.com/frejon93151/netadmin_go/internal/app/get"
	"github.com/frejon93151/netadmin_go/internal/app/models"
)

var Get = &getStruct{
	Devices:                 get.Devices,
	DevicePhysicalInterface: get.DevicePhysicalInterfaces,
	PhysicalInterfaces:      get.PhysicalInterfaces,
}

func NewFilterResp[T comparable]() *models.FilterResp[T] {
	return &models.FilterResp[T]{}
}

type getStruct struct {
	Devices                 func(models.DeviceGetOpts) (*http.Response, error)
	DevicePhysicalInterface func(int, *int, *int) (*http.Response, error)
	PhysicalInterfaces      func(models.PhysGetOpts) (*http.Response, error)
}

type modelsStruct struct {
}
