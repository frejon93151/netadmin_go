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
	"github.com/frejon93151/netadmin_go/internal/app/utils"
)

var Get = getStruct{
	Devices: func(opts DeviceGetOpts) (resp *http.Response, err error) {
		newOpts := utils.MapStruct[models.DeviceGetOpts](opts)
		resp, err = get.Devices(newOpts)
		return
	},
	DevicePhysicalInterface: get.DevicePhysicalInterfaces,
	PhysicalInterfaces: func(opts PhysGetOpts) (resp *http.Response, err error) {
		newOpts := utils.MapStruct[models.PhysGetOpts](opts)
		resp, err = get.PhysicalInterfaces(newOpts)
		return
	},
}

var Post = postStruct{
	DeviceClone: func(id int, opts DevicePostOpts) (resp *http.Response, err error) {
		newOpts := utils.MapStruct[models.DevicePostOpts](opts)
		resp, err = post.DeviceClone(id, newOpts)
		return
	},
}

var Put = putStruct{
	Device: func(id int, opts DevicePostOpts) (resp *http.Response, err error) {
		newOpts := utils.MapStruct[models.DevicePostOpts](opts)
		resp, err = put.Device(id, newOpts)
		return
	},
}

func NewFilterResp[T comparable]() FilterResp[T] {
	return FilterResp[T]{}
}

type getStruct struct {
	Devices                 func(DeviceGetOpts) (*http.Response, error)
	DevicePhysicalInterface func(int, *int, *int) (*http.Response, error)
	PhysicalInterfaces      func(PhysGetOpts) (*http.Response, error)
}

type postStruct struct {
	DeviceClone func(int, DevicePostOpts) (*http.Response, error)
}

type putStruct struct {
	Device func(int, DevicePostOpts) (*http.Response, error)
}

type FilterResp[T comparable] models.FilterResp[T]
type DeviceDTO models.DeviceDTO
type DeviceGetOpts models.DeviceGetOpts
type DevicePostOpts models.DevicePostOpts
type PhysDTO models.PhysDTO
type PhysGetOpts models.PhysGetOpts
