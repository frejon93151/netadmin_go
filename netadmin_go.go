/*
Copyright © 2024 HALMSTADS STADSNÄT AB fredrik.jonsson1@halmstad.se
*/
package netadmin_go

import (
	"net/http"

	"github.com/frejon93151/netadmin_go/internal/app/get"
	"github.com/frejon93151/netadmin_go/internal/app/models"
	"github.com/frejon93151/netadmin_go/internal/app/post"
	"github.com/frejon93151/netadmin_go/internal/app/put"
)

var Get = &getStruct{
	Devices:                 get.Devices,
	DevicePhysicalInterface: get.DevicePhysicalInterfaces,
	PhysicalInterfaces:      get.PhysicalInterfaces,
}

var Post = &postStruct{
	DeviceClone: post.DeviceClone,
}

var Put = &putStruct{
	Device: put.Device,
}

func NewFilterResp[T comparable]() *models.FilterResp[T] {
	return &models.FilterResp[T]{}
}

var Models = &modelsStruct{
	DeviceDTO:      models.DeviceDTO{},
	DeviceGetOpts:  models.DeviceGetOpts{},
	DevicePostOpts: models.DevicePostOpts{},
	PhysDTO:        models.PhysDTO{},
	PhysGetOpts:    models.PhysGetOpts{},
}

type getStruct struct {
	Devices                 func(models.DeviceGetOpts) (*http.Response, error)
	DevicePhysicalInterface func(int, *int, *int) (*http.Response, error)
	PhysicalInterfaces      func(models.PhysGetOpts) (*http.Response, error)
}

type postStruct struct {
	DeviceClone func(int, models.DevicePostOpts) (*http.Response, error)
}

type putStruct struct {
	Device func(int, models.DevicePostOpts) (*http.Response, error)
}

type modelsStruct struct {
	DeviceDTO      models.DeviceDTO
	DeviceGetOpts  models.DeviceGetOpts
	DevicePostOpts models.DevicePostOpts
	PhysDTO        models.PhysDTO
	PhysGetOpts    models.PhysGetOpts
}
