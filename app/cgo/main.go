package main

/* cgo 的基本步骤
1. 如果使用外部的库，需要在编译时链接外部库，否则不需要
2. 通过注释的方式，引入C语言的头文件，或者直接写extern xxx，并紧挨着写import "C"
3. 直接在注释中写C的具体实现，这样不需要额外的库，也不需要额外的c文件
4. 包中的c文件是可以自动编译的，当然需要CGO_ENABLED=1
*/

/*
extern int simple(void);
*/
import "C"

import (
	"cgo/sdk"
)

func main() {
	sdk.Say()
	C.simple()
}
