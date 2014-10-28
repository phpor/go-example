package main

// 参考: https://gobyexample.com/json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type Response1 struct { // 字段之间换行(或分号)分隔而不是逗号分隔，相同类型的N个字段（如果为了省略书写前者的类型，可以逗号分隔）
	Page   int
	Fruits []string
}
type Response2 struct { // 关于最后一个字符串的语法定义参考： http://golang.org/ref/spec#StructType
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

//func main() {
//	decodeXx()
//}

func main() {
	byt := []byte(`{"ss":["a","b"]}`)
	var dat map[string][]string
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Printf("%T\t %v\n", dat, dat)
	ss := dat["ss"]
	fmt.Printf("%T\t%v\n", dat["ss"], dat["ss"])
	fmt.Printf("%T\t%v\n", ss, ss)
}

func decodeArray() {
	doc := []byte(`["_2AkMkqhVta8N1rAFXnP8RzWvqboxH-jyXfg0bAn7oJxImHR19hmlmraB6tSNZoC2qBSA9aq331ynZaJMW","sub exipred"]`)
	result := []string{}

	json.Unmarshal(doc, &result)
	fmt.Printf("%+v", result)
}


func decodeObject() {
	//关于 json.Marshal 的说明参看源码中的注释，有较详细的说明（如：非法utf-8字符的处理办法）
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Printf("%T: %v\n", dat["num"], dat["num"])
	fmt.Printf("%T: %v\n", num, num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	// encode 到变量中
	var bw bytes.Buffer
	enc = json.NewEncoder(&bw)
	enc.Encode(d)

	var out map[string]int
	dec := json.NewDecoder(&bw)
	dec.Decode(&out)

	fmt.Printf(bw.String())
	fmt.Println(out)

	var out2 interface{}
	dec = json.NewDecoder(&bw)
	dec.Decode(&out2)

	fmt.Println(out2)
}
