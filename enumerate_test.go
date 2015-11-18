// +build linux

package udev

import (
	"fmt"
	"runtime"
	"testing"
)

func TestEnumerateDeviceSyspaths(t *testing.T) {
	u := Udev{}
	e := u.NewEnumerate()
	dsp, err := e.DeviceSyspaths()
	if err != nil {
		t.Fail()
	}
	if len(dsp) <= 0 {
		t.Fail()
	}
}

func TestEnumerateSubsystemSyspaths(t *testing.T) {
	u := Udev{}
	e := u.NewEnumerate()
	ssp, err := e.SubsystemSyspaths()
	if err != nil {
		t.Fail()
	}
	if len(ssp) == 0 {
		t.Fail()
	}
}

func TestEnumerateDevicesWithFilter(t *testing.T) {
	u := Udev{}
	e := u.NewEnumerate()
	e.AddMatchSubsystem("block")
	e.AddMatchIsInitialized()
	e.AddNomatchSubsystem("mem")
	e.AddMatchProperty("ID_TYPE", "disk")
	e.AddMatchSysattr("partition", "1")
	e.AddMatchTag("systemd")
	//	e.AddMatchProperty("DEVTYPE", "partition")
	ds, err := e.Devices()
	if err != nil || len(ds) == 0 {
		fmt.Println(len(ds))
		t.Fail()
	}
	for i := range ds {
		if ds[i].Subsystem() != "block" {
			t.Error("Wrong subsystem")
		}
		if !ds[i].IsInitialized() {
			t.Error("Not initialized")
		}
		if ds[i].PropertyValue("ID_TYPE") != "disk" {
			t.Error("Wrong ID_TYPE")
		}
		if ds[i].SysattrValue("partition") != "1" {
			t.Error("Wrong partition")
		}
		if !ds[i].HasTag("systemd") {
			t.Error("Not tagged")
		}

		parent := ds[i].Parent()
		parent2 := ds[i].ParentWithSubsystemDevtype("block", "disk")
		if parent.Syspath() != parent2.Syspath() {
			t.Error("Parent syspaths don't match")
		}

	}
}

func TestEnumerateGC(t *testing.T) {
	runtime.GC()
}
