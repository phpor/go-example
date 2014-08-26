package main

import (
	"code.google.com/p/rsc/qr"
	"os"
)

func main() {
	c, e := qr.Encode("text", qr.L)
	if e != nil {
		println(e); return
	}
	fname := "d:\\temp\\qrcode.png"
	f, err := os.Create(fname)
	defer f.Close()
	if err != nil {
		println(err); return
	}
	f.Write(c.PNG())
}

