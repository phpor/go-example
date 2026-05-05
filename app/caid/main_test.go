package main

import (
	"github.com/phpor/go-example/app/caid/caid-sdk"
	"testing"
)

func Test_buildDeviceInfo(t *testing.T) {
	tests := []struct {
		name string
		want caid.DeviceInfo
	}{
		{name:"test1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getCaid()
		})
	}
}

func Test_GetCaid(t *testing.T) {
	deviceInfo := getDeviceInfo()
	t.Log(GetCaid(deviceInfo))
	t.Log(GenerateCaid(deviceInfo))
}

func BenchmarkBenchMark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BenchMark()
	}
}

func getDeviceInfo() *DeviceInfo {
	return &DeviceInfo{
		BootTimeInSec: "1358889920",
		CountryCode:   "CN",
		Language:      "zh_hans_CN",
		DeviceName:    "e910dddb2748c36b47fcde5dd720eec1",
		SystemVersion: "13.1.1",
		Machine:       "iPhone10,3",
		CarrierInfo:   "中国移动",
		Memory:        "4047224832",
		Disk:          "127938088960",
		SysFileTime:   "1595214620.383940",
		Model:         "D22AP",
		TimeZone:      "28800",
	}
}

func BenchmarkGetCaid(b *testing.B) {
	var succ, fail int
	for i := 0; i < b.N; i++ {   // 450us/op
		deviceInfo := &DeviceInfo{
			BootTimeInSec: "1358889920",
			CountryCode:   "CN",
			Language:      "zh_hans_CN",
			DeviceName:    "e910dddb2748c36b47fcde5dd720eec1",
			SystemVersion: "13.1.1",
			Machine:       "iPhone10,3",
			CarrierInfo:   "中国移动",
			Memory:        "4047224832",
			Disk:          "127938088960",
			SysFileTime:   "1595214620.383940",
			Model:         "D22AP",
			TimeZone:      "28800",
		}
		ret := GetCaid(deviceInfo)
		if len(ret) > 0 {
			succ++
		} else {
			fail++
		}
	}
	b.Logf("succ: %d, fail: %d", succ, fail)
}