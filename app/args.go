package main

import "sync"

//可变参数的用法
func main() {
	strs := []string{"a", "b", "c", "d", "e"}
	var wg sync.WaitGroup
	for _, v := range strs {
		a := map[string]string{} // 这里的 a 在每次循环都是一个新的变量，所以，构造完给一个新的协程是没问题的
		a[v] = v
		wg.Add(1)
		go func() {
			for _, j := range a {
				println(j)
			}
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			println(v) // 这里的v总是最后一个，因为v在循环中只有一个变量
			wg.Done()
		}()
	}
	wg.Wait()
	// append 定义： func append(slice []Type, elems ...Type) []Type

	//b := append([]byte("abc"), []byte{68, 69}...)  // ok
	//b := append([]byte("abc"), byte(67), []byte{68, 69}...)  // not ok
	//b := append([]byte("abc"), byte(67), byte(68))  // ok
}
