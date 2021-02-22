package main

// 关于 c 能不能调 Go
import (
	"fmt"
	sdk_by_go "github.com/phpor/go-example/app/c-call-go/sdk-by-go"
)

func main() {
	fmt.Printf("1 + 2 = %d\n", sdk_by_go.MyAdd(1, 2))
}
