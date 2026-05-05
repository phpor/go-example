package main

import (
	"encoding/base64"
	"regexp"
)

const (
	osFlagAndroid = 0x01
	osFlagIos = 0x02

	osAndroid = "android"
	osIos = "ios"
	osUnknown = "unknown"
)
var (
	osVersionExp = regexp.MustCompile("^[0-9][0-9.]$")
	base64Encode      = base64.URLEncoding.WithPadding('.')
)

func main() {
	println(parseOsFromTid("01Ag1JODNZoWSZqiz5Z4gxEr28"))
	println(parseOsFromTid("01AgZVzqVdifTjRpxlQJZDGadw"))
}

func parseOsFromTid(tid string) string {
	name := osUnknown
	if len(tid) < 5 {
		return name
	}
	buf, err := base64Encode.DecodeString(tid[2:])
	if err != nil {
		return name
	}
	if len(buf) < 4 {
		return name
	}
	if buf[0] != 0x02 {
		return name
	}
	switch buf[1] & 0x03 {
	case osFlagAndroid:
		return osAndroid
	case osFlagIos:
		return osIos
	}
	return name
}
