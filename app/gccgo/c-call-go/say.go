package c_call_go

//export say
func Say() {
	println("hello world")
}
