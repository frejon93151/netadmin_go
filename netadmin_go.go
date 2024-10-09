/*
Copyright © 2024 HALMSTADS STADSNÄT AB fredrik.jonsson1@halmstad.se
*/
package netadmin_go

import (
	"net/http"

	"github.com/frejon93151/netadmin_go/internal/app/Models"
	"github.com/frejon93151/netadmin_go/internal/app/get"
	"github.com/frejon93151/netadmin_go/internal/app/post"
	"github.com/frejon93151/netadmin_go/internal/app/put"
)

var Get = getStruct{
	Devices:                 get.Devices,
	DevicePhysicalInterface: get.DevicePhysicalInterfaces,
	PhysicalInterfaces:      get.PhysicalInterfaces,
}

var Post = postStruct{
	DeviceClone: post.DeviceClone,
}

var Put = putStruct{
	Device: put.Device,
}

func NewFilterResp[T comparable]() Models.FilterResp[T] {
	return Models.FilterResp[T]{}
}

type getStruct struct {
	Devices                 func(Models.DeviceGetOpts) (*http.Response, error)
	DevicePhysicalInterface func(int, *int, *int) (*http.Response, error)
	PhysicalInterfaces      func(Models.PhysGetOpts) (*http.Response, error)
}

type postStruct struct {
	DeviceClone func(int, Models.DevicePostOpts) (*http.Response, error)
}

type putStruct struct {
	Device func(int, Models.DevicePostOpts) (*http.Response, error)
}
