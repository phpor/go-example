package main

//可变参数的用法
func main() {
  // append 定义： func append(slice []Type, elems ...Type) []Type
  
  b := append([]byte("abc"), []byte{68, 69}...)  // ok
  b := append([]byte("abc"), byte(67), []byte{68, 69}...)  // not ok
  b := append([]byte("abc"), byte(67), byte(68))  // ok
}
