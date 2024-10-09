/*
Copyright © 2024 HALMSTADS STADSNÄT AB fredrik.jonsson1@halmstad.se
*/
package netadmin_go

import (
	"net/http"

	Models1 "github.com/frejon93151/netadmin_go/internal/app/Models"
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

func NewFilterResp[T comparable]() Models1.FilterResp[T] {
	return Models1.FilterResp[T]{}
}

type Models interface {
	Models1.DeviceDTO
	Models1.DeviceGetOpts
	Models1.DevicePostOpts
	Models1.PhysDTO
	Models1.PhysGetOpts
}

type getStruct struct {
	Devices                 func(Models1.DeviceGetOpts) (*http.Response, error)
	DevicePhysicalInterface func(int, *int, *int) (*http.Response, error)
	PhysicalInterfaces      func(Models1.PhysGetOpts) (*http.Response, error)
}

type postStruct struct {
	DeviceClone func(int, Models1.DevicePostOpts) (*http.Response, error)
}

type putStruct struct {
	Device func(int, Models1.DevicePostOpts) (*http.Response, error)
}

type modelsStruct struct {
	DeviceDTO      Models1.DeviceDTO
	DeviceGetOpts  Models1.DeviceGetOpts
	DevicePostOpts Models1.DevicePostOpts
	PhysDTO        Models1.PhysDTO
	PhysGetOpts    Models1.PhysGetOpts
}
