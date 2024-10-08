package main

import (
	"net/http"

	"github.com/frejon93151/netadmin_go/internal/app/get"
	"github.com/frejon93151/netadmin_go/internal/app/models"
)

var Get = &getType{
	Devices:                 get.Devices,
	DevicePhysicalInterface: get.DevicePhysicalInterfaces,
	PhysicalInterfaces:      get.PhysicalInterfaces,
}

type getType struct {
	Devices                 func(models.DeviceGetOpts) (*http.Response, error)
	DevicePhysicalInterface func(int, *int, *int) (*http.Response, error)
	PhysicalInterfaces      func(models.PhysGetOpts) (*http.Response, error)
}
