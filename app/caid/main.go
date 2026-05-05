package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/phpor/go-example/app/caid/caid-sdk"
	"strings"
)

type DeviceInfo struct {   //计算caid需要的信息， 共 12 项
	BootTimeInSec string   //设备启动时间
	CountryCode string     //国家
	Language string        //语言
	DeviceName string      //设备名称,  不允许为空，为空得不到caid
	SystemVersion string   //系统版本
	Machine string         //设备 machine
	CarrierInfo string     //运营商
	Memory string          //物理内存
	Disk string            //磁盘容量
	SysFileTime string     //系统更新时间
	Model string           //设备model
	TimeZone string        //时区
}

type CaidInfo map[string]string //生成的caid的结果
func (cr *CaidInfo) String() string {
	return fmt.Sprintf("%v", *cr)
}

func main() {
	getCaid()
}

func GenerateCaid(info *DeviceInfo) CaidInfo {
	//"devicename|1595214620.383940|D22AP|iPhone10,3|2bae66d0fabaf3608c3aab2da56dcd5f|9F1919DD414AA74B105972E01ECA84CD"

	key1 := "2bae66d0fabaf3608c3aab2da56dcd5f"
	key2 := "9F1919DD414AA74B105972E01ECA84CD"
	ss := []string{
		info.DeviceName, info.SysFileTime, info.Model, info.Machine, key1, key2,
	}
	str := strings.Join(ss, "|")
	m := md5.Sum([]byte(str))

	caidInfo := CaidInfo{}
	caidInfo["20201230"] = hex.EncodeToString(m[:])
	return caidInfo
}


func GetCaid(info *DeviceInfo) CaidInfo {
	di := caid.NewDeviceInfo()
	defer caid.DeleteDeviceInfo(di)
	di.SetBoot_time_in_sec(info.BootTimeInSec)
	di.SetCountry_code(info.CountryCode)
	di.SetLanguage(info.Language)
	di.SetDevice_name(info.DeviceName)
	di.SetSystem_version(info.SystemVersion)
	di.SetMachine(info.Machine)
	di.SetCarrier_info(info.CarrierInfo)
	di.SetMemory(info.Memory)
	di.SetDisk(info.Disk)
	di.SetSys_file_time(info.SysFileTime)
	di.SetModel(info.Model)
	di.SetTime_zone(info.TimeZone)
	ret := caid.GenerateCAID(di)
	defer caid.DeleteCaidVector(ret)
	caidInfo := CaidInfo{}
	for i := 0; i < int(ret.Size()) ; i++  {
		pair := ret.Get(i)
		caidInfo[pair.GetFirst()] = pair.GetSecond()
	}
	return caidInfo
}

func getCaid() {
	d := buildDeviceInfo()
	ret := caid.GenerateCAID(d)
	fmt.Printf("%#v\n", ret.Size())
	for i := 0; i < int(ret.Size()) ; i++  {
		pair := ret.Get(i)
		fmt.Printf("%#v, %#v\n", pair.GetFirst(), pair.GetSecond())
	}
}

func buildDeviceInfo() caid.DeviceInfo{
	d := caid.NewDeviceInfo()
	d.SetBoot_time_in_sec("1358889920")
	d.SetCountry_code("CN")
	d.SetLanguage("zh_hans_CN")
	d.SetDevice_name("e910dddb2748c36b47fcde5dd720eec1")
	d.SetSystem_version("13.1.1")
	d.SetMachine("iPhone10,3")
	d.SetCarrier_info("aaaa")
	d.SetMemory("4047224832")
	d.SetDisk("127938088960")
	d.SetSys_file_time("1595214620.383940")
	d.SetModel("D22AP")
	d.SetTime_zone("28800")
	return d
}

func BenchMark() bool {
	d := buildDeviceInfo()
	ret := caid.GenerateCAID(d)
	caid.DeleteDeviceInfo(d)
	defer func() {
		caid.DeleteCaidVector(ret) // 删除容器的时候回自动删除容器内部的数据的
	}()
	if ret.Size() > 0 {
		//for i := 0; i < int(ret.Size()) ; i++  {
		//	caid.DeleteCaidMap(ret.Get(i))
		//}
		return true
	}
	return false
}
