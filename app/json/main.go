package main

import (
	"encoding/json"
	"fmt"
)

type person struct { // 表面上json包没法new一个person值，但是，反射可以new这个结构，这里的私有结构体的私有只是提示给编译器看的
	Name string
	Age  int
}

type Team struct {
	Leader   *person // 指针也能被填充
	Follower person
}

func main() {
	testEscape()
}

func testMap() {
	//s := []byte(`{"a":"A", "b":2}`)
	s := []byte(`{"mfp":"{\"1\":\"13.7\",\"11\":\"E99163D8-5A2B-47ED-A9F3-55555AE5E5E2\",\"18\":\"iPhone\",\"2\":\"796997f731e5681f6e9e5ea298866da43e1cf2f0\",\"23\":\"zh-Hans-CN\",\"24\":\"460,02,中国移动\",\"33\":\"1614525706\",\"37\":\"45.276970\",\"38\":\"3117580288\",\6002\",\"69\":\"iPhone10,2\",\"70\":\"255989469184\",\"73\":\"1597084210.236371\",\"75\":\"CN\",\"80\":\"L3Zhci9tb2JpbGUvTGlicmFyeS9Vc2VyQ29uZmlndXJhdGlvblByb2ZpbGVzL1B1YmxpY0luZm8vTUNNZXRhLnBsaXN0\",\"83\":\"none\",\"85\":\"D21AP\"}","src":"getaid"}`)
	m := map[string]interface{}{}
	err := json.Unmarshal(s, &m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", m)
}

func testStruct() {
	str := `{"leader":{"Name":"wang", "age":22},"Follower":{"Name":"sun", "age":20}}`
	t := &Team{}
	if err := json.Unmarshal([]byte(str), t); err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("leader: %#v, follower: %#v", t.Leader, t.Follower)

}

func testEscape() {
	a := struct {
		A string
	}{`{"1":"13.6","11":"E99163D8-5A2B-47ED-A9F3-55555AE5E5E2","18":"张小花的 iPhone","2":"696997f731e5681f6e9e5ea298866da43e1cf2f0","23":"zh-Hans-CN","24":"460,02,中国移动","33":"1614525706","37":"45.276970","38":"3117580288","4":"46002","69":"iPhone10,2","70":"255989469":"1597084210.236371","75":"CN","80":"L3Zhci9tb2JpbGUvTGlicmFyeS9Vc2VyQ29uZmlndXJhdGlvblByb2ZpbGVzL1B1YmxpY0luZm8vTUNNZXRhLnBsaXN0","83":"none","85":"D21AP"}`}
	b, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(b))
}
