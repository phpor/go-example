// 参考资料： http://cjmxp007.blog.163.com/blog/static/35473837201231115825658/

package main

import (
	"reflect"
	"fmt"
)

func main() {
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

