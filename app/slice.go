// 参考资料： http://www.tuicool.com/articles/rYVf2e	（写的比较到位）
package main

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func slice_is_readonly(ss []string) {
	fmt.Printf("ss in func: %p\n", ss) // slice 通过参数传递时是传地址的
}
func main() {
	subSlice()
}

func subSlice() {
	a := []byte("abc")
	fmt.Printf("%s", a[2:2]) // 前后相同就是没有,前闭后开
	//testSliceFilter()

	println("test2:")
	b := []int{1}
	println(b[1:]) // 虽然这里的1是越界的，但是，截取不会报错
	//println(b[1])  //虽然这个看起来和上面的切片截取差不多，但是，这里就是越界错误。

	println("test 3")
	c := []int{1, 2, 3, 4, 5, 6, 7}
	c1 := c[2:4] // 这里的c1 的cap是从2到c的cap结束，这里c1 长度2 ，cap 5
	println("c1: ", c1)
	println("append...")
	_ = append(c1, 8) // 这里没有修改c1的长度，所以，c1的内容没变，但是，因为cap允许就没有发生内存移位，于是，影响到了c
	println("c:", c)
	println("c1: ", c1)
	fmt.Printf("c: %#v\n", c)
	fmt.Printf("c1: %#v\n", c1)

}

func testSliceFilter() {
	ss := []string{"phpor", "zhangsan", "php", "go", "python", "c"}
	SliceFilter(&ss, func(i int) bool {
		fmt.Printf("%v\n", ss[i])
		return strings.HasPrefix(ss[i], "p")
	})
	fmt.Printf("%v\n", ss)
}

func testSlice3() {
	ss := make([]string, 2, 3)
	ss[0] = "aa"
	ss[1] = "bb"
	modifySs := func(ss []string) {
		ss[0] = "AA"
		ss = append(ss, "cc") // 会产生新的ss， 不是因为原来的ss容量不够，而是，这根本就是一个新的ss
	}
	modifySs(ss)
	fmt.Printf("%v", ss) // [AA bb]
	// 这说明: slice 不能算是传地址的，说成引用传递比较好，修改其内容是可以的，但是修改其本身是不行的

}
func testSlice2() {
	a := make([]byte, 18)
	a[0] = 1
	binary.LittleEndian.PutUint32(a[1:5], uint32(2))
	copy(a[5:], []byte{3})
	println(a)
}

func testSlice1() {

	s := "hello_world"
	fmt.Println(s[:5]) //字符串默认可以当slice使用
	fmt.Println(s[5:6])
	fmt.Println(s[6:])

	ss := []string{"hello", "word"}
	fmt.Printf("ss in main: %p\n", ss)
	fmt.Printf("ss[1:] in main: %p\n", ss[1:])
	slice_is_readonly(ss[1:])

	sss := []string{"你好", "世界"}
	fmt.Println(sss[1:])
	fmt.Println(sss[:1])
	fmt.Println("golang 默认没有字符集转换类库，需要第三方类库")

	// 排序的示例
	ssss := []string{"c2", "c1", "b1", "b2", "a1", "a2"}
	fmt.Printf("before sort: %p\n", ssss)
	sort.Strings(ssss)
	fmt.Printf("after sort: %p\n", ssss)
	fmt.Println(ssss)
	fmt.Println(ssss[1])
	ssss[1] = "A2"
	fmt.Printf("%T :%p :%s\n", ssss, ssss, ssss) //slice 是可以被修改的

	println("slice or array:")
	arr := [2][2]int{{11, 12}, {21, 22}} // 课本上说，这个是数组，不是切片，但是很多时候用起来和切片一样一样的
	arr2 := [...]int{1, 2, 3}            // 这是个数组
	s6 := []int{1, 2, 3}                 // 这是个切片
	//var s7 []int = [...]int{1,2,3}		  // 这个赋值是错误的，类型不同，生命的是切片，赋值的是数组
	fmt.Printf("%T : %v :%v\n", arr, arr, arr[1])
	fmt.Printf("%T : %v :%v\n", arr2, arr2, arr2[1])
	fmt.Printf("%T : %v :%v\n", s6, s6, s6[1])

	println("append to slice:")
	//slice 的追加
	s1 := []string{"a", "b"}
	s2 := []string{"c2", "d"}
	fmt.Printf("&s1: %p\n", s1)

	// 一般会这么写：
	s1 = append(s1, s2...)
	fmt.Printf("%p: %v \tcap(s1):%d\n", s1, s1, cap(s1)) // s1的地址可发生变化，也可能不变化，主要取决于s1的剩余容量够不够放下s2

	s3 := append(s1, s2...) // 注意：追加的是slice中的元素，而不是slice，所以这里用 ... 语法，将slice结构成多个元素参数
	fmt.Printf("%p: %v\n%p: %v\n", s1, s1, s3, s3)
}

func SliceFilter(data interface{}, f func(int) bool) {
	if reflect.TypeOf(data).Kind() != reflect.Ptr || reflect.ValueOf(data).Elem().Kind() != reflect.Slice {
		panic("data must be slice pointer")
	}

	va := reflect.ValueOf(data).Elem()

	j := 0
	l := va.Len()
	for i := 0; i < l; i++ {
		if f(i) {
			j++
			continue
		}

		// 从最后面找一个填进来
		for {
			if i+1 >= l { // 如果 i 已经是最后一个了，则退出循环
				break
			}
			l-- // 最后一个要么移动到前面，要么不符合条件被丢弃，所以这里直接l--
			if f(l) {
				va.Index(j).Set(va.Index(l))
				j++
				break
			}
		}
	}

	va.SetLen(j)
}
