package main

// 这个需要gcc才能编译，而且能否编译成功还要看看目标平台结构和C库本身（32位？64位）是否一致
// 参考使用go-charset
import (
	iconv "github.com/djimenez/iconv-go"
)

func main() {
	iconv.ConvertString("我是UTF-8", "utf-8", "gbk")
}
