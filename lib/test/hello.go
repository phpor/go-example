package test

/**
1. 通过 go install 来安装编译后的包
2. go 目前还无法做到只发布编译后的文件而不共享源代码
 */
import "fmt"

func Hello(word string) {
	fmt.Println(word)
}
