package main

import (
	"fmt"
	"git.intra.weibo.com/gopkg/gx-caid"
)

func main() {
	// 生成caid

	res := caid.GenerateCaid(buildDeviceInfo())
	fmt.Printf("%+v", res.GetAll())

}

func buildDeviceInfo() *caid.DeviceInfo {
	return &caid.DeviceInfo{
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