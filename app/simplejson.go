package main

// simplejson 确实用起来simple多了
import (
	json "github.com/bitly/go-simplejson"
	jj "encoding/json"
	"fmt"
	"os"
)

func main() {
	test1()
}

func test1() {
	text := []byte(`{
				"int_array":[1,2,3],
				"map":{
					"string_array": ["v1", "v2"]
				},
				"int": 2,
				"bool": true
			}`)
	j, err := json.NewJson(text)
	checkError(err)
	arr, err := j.Get("int_array").Array()
	checkError(err)
	fmt.Println("get int_array:")
	for _, v := range arr {
		fmt.Printf("%s ", v.(jj.Number))    // 这是多么不漂亮的写法啊
	}
	fmt.Println()

	// 使用 GetPath + Must* 还是非常方便的
	fmt.Println("use GetPath & MustArray to get string_array:")
	for _, v := range j.GetPath("map", "string_array").MustArray() {
		fmt.Printf("%s ", v.(string))
	}
	fmt.Println()

	fmt.Println("use GetPath & MustArray to get string_array (no error):")
	for _, v := range j.GetPath("no_exists", "string_array").MustArray() {    // 不存在的信息也不会出错
		fmt.Printf("%s ", v.(string))
	}
	fmt.Println()

	fmt.Printf("int: %d\n", j.Get("int").MustInt())
	fmt.Printf("int: %d\n", j.Get("int_not_exists").MustInt(9)) // 当想要的字段不存在时，使用默认值是不错的
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

