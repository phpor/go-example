// 参考资料： http://cjmxp007.blog.163.com/blog/static/35473837201231115825658/

package main

import (
	"errors"
	"fmt"
	"reflect"
)

type C struct {
	Name string
}
type B struct {
	C *C
}
type A struct {
	B B
}
type II int32

func main() {
	ParseTypeAlias()
	//AccessPtrStruct()
	//DeepReflect()

	//s := struct {
	//	A A
	//	A2 string
	//	A3 *string
	//}{A:A{B{}}, A2: "haha"}
	//fmt.Printf("%v\n", s)
	//b,_ := json.Marshal(deepInspectStruct(s, 10))
	//fmt.Printf("%s", b)
}

func ParseTypeAlias() {
	var ii II = 10
	var jj interface{}
	jj = ii
	v := reflect.ValueOf(jj)
	//v := reflect.TypeOf(jj)
	switch f := jj.(type) {
	case int32:
		fmt.Printf("xxx: %T\n", f)
	case II:
		fmt.Printf("xxxII: %T\n", f)

	}
	if v.Kind() == reflect.Int32 {

		fmt.Printf("这里的jj确实是int32，而不是反射出来II\n")
	}
	fmt.Printf("%v ", v.Interface())
}

// 这个深度递归求结构体字段信息没符合预期，另外，这个有可以无限递归下去，所以，就算可以，使用起来也需要谨慎，可以添加一个深度检测
// 遇到结构体的值为nil就没法继续获取相关的结构的信息了，这个有点儿出乎意料
func deepInspectStruct(container interface{}, maxDepth int) map[string]interface{} {
	if maxDepth <= 0 {
		maxDepth = 100
	}
	// 为了能自己调用自己，这里就需要先定义一次，或者使用一个中间变量，总之，不能像定义函数一样初始化时直接递归使用自己
	var deepInspect func(container interface{}, deepCounter int) map[string]interface{}
	deepInspect = func(container interface{}, deepCounter int) map[string]interface{} {
		deepCounter++
		c := reflect.TypeOf(container)
		v := reflect.ValueOf(container)
		fieldNum := c.NumField()
		result := map[string]interface{}{}
		for i := 0; i < fieldNum; i++ {
			field := c.Field(i)
			v2 := v.FieldByName(field.Name)
			kind := v2.Kind()
			if kind == reflect.Ptr || kind == reflect.Struct {
				if field.Type.Kind() == reflect.Ptr {
					v2 = v2.Elem()
				}
				if v2.Kind() == reflect.Struct {
					if deepCounter > maxDepth {
						result[field.Name] = "too deep to stop"
						continue
					}
					result[field.Name] = deepInspect(v2.Interface(), deepCounter)
					continue
				}
			}
			result[field.Name] = field.Type.String()
		}
		return result
	}
	return deepInspect(container, 0)
}

type Person struct {
	Name string
}

func getVal(ctx interface{}, arr []string) (interface{}, error) {
	if len(arr) == 0 {
		return nil, errors.New("field Name is empty")
	}
	val := reflect.ValueOf(ctx)

	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil, errors.New("no found")
		}
		val = val.Elem()
	}
	name := arr[0]
	var v interface{}
	switch val.Kind() {
	case reflect.Struct:
		fval := val.FieldByName(name)
		if !fval.IsValid() {
			return nil, errors.New("no found")
		}
		v = fval.Interface()
	default:
		return nil, errors.New("container type no match")
	}

	if len(arr) == 1 {
		return v, nil
	}
	return getVal(v, arr[1:])
}

func DeepReflect() {

	p := struct {
		Name  string
		Child Person
	}{
		Name: "phpor",
		Child: Person{
			Name: "lizhiyuan",
		},
	}
	q := struct {
		X interface{}
	}{
		X: &p,
	}

	//var v interface{}
	//v = p
	//v2, err := getVal(&p, []string{"Child","Name"})
	v2, err := getVal(q, []string{"X", "Child", "Name"})
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("%v\n", v2)
	x := &v2
	*x = "erzi"
	fmt.Printf("%v", q.X)
}

func AccessPtrStruct() {
	var a interface{}

	a = &Person{Name: "phpor"}

	typ := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	vt := v.Type()
	kind := typ.Kind()
	fmt.Printf("%v\n", kind)
	fmt.Printf("%v\n", vt.Kind())

	if kind == reflect.Ptr {
		v2 := v.Elem()
		v2t := v2.Type()
		fmt.Printf("v2t.Kind(): %v\n", v2t.Kind())

		ps := reflect.ValueOf(v2)
		ps.Kind()
		fmt.Printf("%v\n", ps.Kind())

		fval := v2.FieldByName("Name") //todo: 这里有可能panic，需要注意

		fmt.Printf("%v\n", fval.Interface())
		fval.SetString("gaiguo le")
		//if !fval.IsValid() {         // 未找到对应字段，返回nil，尽量避免报错
		//	return nilExpr, nil
		//}
		//val = fval.Interface()
		fmt.Printf("%v", a)
	}

}

func test3() {
	var err interface{} = nil
	a := reflect.TypeOf(err)
	fmt.Printf("%v\n", a)
	str := "abcd"
	i := 1234
	v1 := reflect.ValueOf(str)
	v2 := reflect.ValueOf(i)
	t1 := reflect.TypeOf(str)
	t2 := reflect.TypeOf(i)

	fmt.Printf("%T\n%+v\n\n", v1, v1)
	fmt.Printf("%T\n%+v\n\n", v2, v2)
	fmt.Printf("%T\n%+v\n\n", t1, t1)
	fmt.Printf("%T\n%+v\n\n", t2, t2)
}
