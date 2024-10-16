/*
Copyright © 2024 HALMSTADS STADSNÄT AB fredrik.jonsson1@halmstad.se
*/
package netadmin_go

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/frejon93151/netadmin_go/internal/app/netadminsql"
	"github.com/frejon93151/netadmin_go/internal/app/utils"
)

var (
	filename         = "test.env"
	fatalStr         = "\nfailed at: \nfunc %s = %v %v\nparams: %q\n"
	testDevice       = 31467
	copyOfTestDevice = 33693
	testSite         = 2447
	testAddress      = 51634
)

func loadEnv(t *testing.T) {
	buf, err := utils.ReadFile(filename)
	if err != nil {
		t.Fatalf(fatalStr, utils.NameOf(utils.ReadFile), buf, err, filename)
	}
	env, err := utils.ParseEnv(buf)
	if err != nil {
		t.Fatalf(fatalStr, utils.NameOf(utils.ParseEnv), env, err, buf)
	}
	for key := range env {
		os.Setenv(key, env[key])
	}
}

// TestGetVLANRange tests sql
func TestGetVLANRange(t *testing.T) {
	loadEnv(t)

	vlan, err := netadminsql.GetVLANRange("pw-63:____%")
	if err != nil {
		t.Fatalf(fatalStr, utils.NameOf(netadminsql.GetVLANRange), vlan, err, "pw-63:____%")
	}
	for _, item := range vlan {
		t.Log(item)
	}
}

// TestGetDevice calls Get.Device and checks for a success status code and valid return value
func TestGetDevice(t *testing.T) {
	loadEnv(t)

	var ids []int
	ids = append(ids, testDevice)
	opts := DeviceGetOpts{Ids: ids, PageIndex: 0, ItemsPerPage: 30}
	resp, err := Get.Devices(opts)
	if err != nil {
		t.Fatalf(fatalStr, utils.NameOf(Get.Devices), resp, err, ids)
	}

	defer resp.Body.Close()
	content := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&content)
	if err != nil {
		t.Fatalf(fatalStr, utils.NameOf(json.NewDecoder), nil, err, resp.Body)
	}
	if resp.StatusCode > 299 {
		t.Fatalf(fatalStr, utils.NameOf(Post.DeviceClone), content, testDevice, opts)
	}
	for key, item := range content {
		t.Log(key, ": ", item)
	}
}

// TestGetDeviceAuthentications calls Get.DeviceExtensions and check for a success status code and a valid return value
func TestGetDeviceAuthentications(t *testing.T) {
	loadEnv(t)
	resp, err := Get.DeviceAuthentications(testDevice)
	if err != nil {
		t.Fatalf(fatalStr, utils.NameOf(Get.DeviceAuthentications), resp, err, testDevice)
	}

	defer resp.Body.Close()
	buf, err := io.ReadAll(resp.Body)
	t.Log(string(buf))
	content := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&content)
	if err != nil {
		t.Fatalf(fatalStr, utils.NameOf(json.NewDecoder), content, err, resp.Body)
	}
	if resp.StatusCode > 299 {
		t.Fatalf(fatalStr, utils.NameOf(Get.DeviceAuthentications), content, err, testDevice)
	}
	for key, item := range content {
		t.Log(key, ": ", item)
	}
}

// TestCloneDevice calls Post.DeviceClone and check for a success status code and valid return value
func TestCloneDevice(t *testing.T) {
	loadEnv(t)
	opts := DevicePostOpts{
		Name:                 fmt.Sprintf("COPYOF-%d", testDevice),
		HostName:             fmt.Sprintf("COPYOF-%d", testDevice),
		AdministrativeStatus: "Inaktiv",
	}
	resp, err := Post.DeviceClone(testDevice, opts)
	if err != nil {
		t.Fatalf(fatalStr, utils.NameOf(Post.DeviceClone), resp, err, testDevice, opts)
	}

	defer resp.Body.Close()
	content := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&content)
	if err != nil {
		t.Fatalf(fatalStr, utils.NameOf(json.NewDecoder), nil, err, resp.Body)
	}
	if resp.StatusCode > 299 {
		t.Fatalf(fatalStr, utils.NameOf(Post.DeviceClone), content, testDevice, opts)
	}
	for key, item := range content {
		t.Log(key, ": ", item)
	}
}

// TestUpdateDevice calls Put.Device and check for a success status code and valid return value
func TestUpdateDevice(t *testing.T) {
	loadEnv(t)

	opts := DevicePostOpts{
		Name:                 fmt.Sprintf("COPYOF-%d", testDevice),
		HostName:             fmt.Sprintf("COPYOF-%d", testDevice),
		AdministrativeStatus: "Inaktiv",
		SiteId:               testSite,
		ManagementAddress:    "0.0.0.0",
		SerialNumber:         "0x0x0x0x0x0",
		ReferenceName:        "COPY OF sw-hsab-lab",
	}
	resp, err := Put.Device(copyOfTestDevice, opts)
	if err != nil {
		t.Fatalf(fatalStr, utils.NameOf(Put.Device), resp, err, copyOfTestDevice, opts)
	}

	defer resp.Body.Close()
	content := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&content)
	if err != nil {
		t.Fatalf(fatalStr, utils.NameOf(json.NewDecoder), nil, err, resp.Body)
	}
	if resp.StatusCode > 299 {
		t.Fatalf(fatalStr, utils.NameOf(Post.DeviceClone), content, testDevice, opts)
	}
	for key, item := range content {
		t.Log(key, ": ", item)
	}
}

// TestGetDevicePhysIf calls Get.DevicePhysicalInterfaces
func TestGetDevicePhysIf(t *testing.T) {
	loadEnv(t)

	resp, err := Get.DevicePhysicalInterface(testDevice, 0, 60)
	if err != nil {
		t.Fatalf(fatalStr, utils.NameOf(Get.DevicePhysicalInterface), resp, err, testDevice, 0, 50)
	}

	defer resp.Body.Close()
	content := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&content)
	if err != nil {
		t.Fatalf(fatalStr, utils.NameOf(json.NewDecoder), nil, err, resp.Body)
	}
	if resp.StatusCode > 299 {
		t.Fatalf(fatalStr, utils.NameOf(Post.DeviceClone), content, testDevice, 0, 50)
	}
	for key, item := range content {
		t.Log(key, ": ", item)
	}
}

// TestPutPhysIf frist calls Get.DevicePhysicalInterface(testDevice) then Put.PhysIf for each device in the return
func TestPutPhysIf(t *testing.T) {
	loadEnv(t)

	physIfRes, err := Get.DevicePhysicalInterface(copyOfTestDevice, 0, 60)
	if err != nil {
		t.Fatalf(fatalStr, utils.NameOf(Get.DevicePhysicalInterface), physIfRes, err, copyOfTestDevice, 0, 60)
	}

	defer physIfRes.Body.Close()
	physIfContent := NewFilterResp[PhysDTO]()
	err = json.NewDecoder(physIfRes.Body).Decode(&physIfContent)
	if err != nil {
		t.Fatalf(fatalStr, utils.NameOf(json.NewDecoder), nil, err, physIfRes.Body)
	}
	if physIfRes.StatusCode > 299 {
		t.Fatalf(fatalStr, utils.NameOf(Get.DevicePhysicalInterface), physIfContent, copyOfTestDevice, 0, 60)
	}
	var physIf []PhysDTO
	for index, physItem := range physIfContent.Items {
		if physItem.Number < 49 {
			t.Log(index, ": ", physItem.Id, " ", physItem.Number)
			physIf = append(physIf, physItem)

			opts := PhysPostOpts{
				Identifier:           physItem.Identifier,
				Option82:             "",
				Panel:                physItem.Panel,
				OperationalStatus:    physItem.OperationalStatus,
				AdministrativeStatus: physItem.AdministrativeStatus,
				Fqpn:                 physItem.Fqpn,
				Description:          physItem.Description,
			}
			resp, err := Put.PhysIf(physItem.Id, opts)
			if err != nil {
				t.Fatalf(fatalStr, utils.NameOf(Put.PhysIf), nil, err, physItem.Id, opts)
			}
			t.Log(resp.Status)
			defer resp.Body.Close()
			content := make(map[string]interface{})
			err = json.NewDecoder(resp.Body).Decode(&content)
			for key, item := range content {
				t.Log(key, ": ", item)
			}
		}
	}
}

func TestPhysIfConnect(t *testing.T) {
	loadEnv(t)

	physIfRes, err := Get.DevicePhysicalInterface(testDevice, 0, 60)
	if err != nil {
		t.Fatalf(fatalStr, utils.NameOf(Get.DevicePhysicalInterface), physIfRes, err, testDevice, 0, 60)
	}

	defer physIfRes.Body.Close()
	physIfContent := NewFilterResp[PhysDTO]()
	err = json.NewDecoder(physIfRes.Body).Decode(&physIfContent)
	if err != nil {
		t.Fatalf(fatalStr, utils.NameOf(json.NewDecoder), nil, err, physIfRes.Body)
	}
	if physIfRes.StatusCode > 299 {
		t.Fatalf(fatalStr, utils.NameOf(Get.DevicePhysicalInterface), physIfContent, testDevice, 0, 60)
	}

	for index, physItem := range physIfContent.Items {
		if physItem.Number == 2 {
			t.Log(index, ": ", physItem.Id, " ", physItem.Number)

			resp, err := Post.PhysIfConnect(physItem.Id, testAddress)
			if err != nil {
				t.Fatalf(fatalStr, utils.NameOf(Put.PhysIf), nil, err, physItem.Id, testAddress)
			}
			t.Log(resp.Status)
			defer resp.Body.Close()
			content := make(map[string]interface{})
			err = json.NewDecoder(resp.Body).Decode(&content)
			for key, item := range content {
				t.Log(key, ": ", item)
			}
		}
	}
}

func TestPhysIfDisconnect(t *testing.T) {
	loadEnv(t)

	physIfRes, err := Get.DevicePhysicalInterface(testDevice, 0, 60)

	if err != nil {
		t.Fatalf(fatalStr, utils.NameOf(Get.DevicePhysicalInterface), physIfRes, err, testDevice, 0, 60)
	}

	defer physIfRes.Body.Close()
	physIfContent := NewFilterResp[PhysDTO]()
	err = json.NewDecoder(physIfRes.Body).Decode(&physIfContent)
	if err != nil {
		t.Fatalf(fatalStr, utils.NameOf(json.NewDecoder), nil, err, physIfRes.Body)
	}
	if physIfRes.StatusCode > 299 {
		t.Fatalf(fatalStr, utils.NameOf(Get.DevicePhysicalInterface), physIfContent, testDevice, 0, 60)
	}

	for index, physItem := range physIfContent.Items {
		if physItem.Number == 2 {
			t.Log(index, ": ", physItem.Id, " ", physItem.Number)

			resp, err := Post.PhysIfDisconnect(physItem.Id)
			if err != nil {
				t.Fatalf(fatalStr, utils.NameOf(Put.PhysIf), nil, err, physItem.Id, testAddress)
			}
			t.Log(resp.Status)
			defer resp.Body.Close()
			content := make(map[string]interface{})
			err = json.NewDecoder(resp.Body).Decode(&content)
			for key, item := range content {
				t.Log(key, ": ", item)
			}
		}
	}
}
