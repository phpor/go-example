package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type Obj struct {
	A  string     `hbase:"a" json:"json_a"`
	B  string     `hbase:"b" json:"json_b"`
	T  *time.Time `hbase:"t"`
	T2 time.Time
}

// Say 可以打印几句话
func (o *Obj) Say() {
	fmt.Printf("say some thing")
}

func main() {
	testInject()
}

func testReflectAndSet() {

	obj := &Obj{
		A: "",
		B: "",
		T: &time.Time{},
	}
	reflectAndSet(obj, "T2", "AAA")
	c, _ := json.Marshal(obj)
	fmt.Printf("%s", c)

	reflectTags(Obj{})
}

func testBytesToInt() {
	s1 := "1111"
	s2 := s1
	*(&s1) = "2222"
	println(s2)

	fmt.Println(len("\x00\x00\x00\x00>\xC4\xD8\x9C"))
	//println(BytesToInt([]byte{1,2,3,4}))
	now := time.Now().Unix()
	println(now)
	//b := IntToBytes(now, 4)
	b := packTime(uint32(now))
	n := BytesToInt(b)
	println(n)
	println(int(n))
	println("--------")

	v := make([]byte, 4)
	binary.BigEndian.PutUint32(v, uint32(now))
	n2 := binary.BigEndian.Uint32(b)
	println(n2)
	println(int(n2))

}

func testMaxInt() {
	println(maxInt())
	println(maxInt(1))
	println(maxInt(1, 2))
	println(maxInt(1, 2, 3, 4))
	println(maxInt(5, 2, 6, 4))
	println(maxInt(7, 2, 6, 4))
}

func packTime(t uint32) []byte {
	v := make([]byte, 4)
	binary.BigEndian.PutUint32(v, t)
	return v
}

func maxInt(i ...int) int {
	length := len(i)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return i[0]
	}
	if i[0] > i[1] {
		i[1] = i[0]
	}
	return maxInt(i[1:]...)
}

func BytesToInt(b []byte) int32 {
	buf := bytes.NewBuffer(b)
	var data int32
	err := binary.Read(buf, binary.BigEndian, &data)
	if err != nil {
		println(err.Error())
	}
	return data
}

func IntToBytes(n int64, len int) []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, n)
	if buf.Len() < len {
		return buf.Bytes()
	}
	return buf.Bytes()[(buf.Len() - len):]
}

func reflectTags(obj Obj) {
	s := reflect.TypeOf(obj)
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		name := f.Tag.Get("hbase")
		fmt.Printf("%s: %s\n", f.Name, name)
	}
}

func reflectAndSet(obj *Obj, fieldName, val string) {
	if obj == nil {
		return
	}
	s := reflect.ValueOf(obj).Elem()
	f := s.FieldByName(fieldName)
	if !f.CanSet() {
		return
	}
	switch f.Kind() {
	case reflect.String:
		f.SetString(val)
	case reflect.Ptr: // 如果指针为空，还不能直接赋值
		if t, ok := f.Interface().(*time.Time); ok {
			println("haha")
			*t = time.Now()
		}
	case reflect.Struct: // 这种方法给结构体赋值还不错
		if _, ok := f.Interface().(time.Time); ok {
			f.Set(reflect.ValueOf(time.Now()))
		}
	default:
		println(f.Kind().String())

	}
}
