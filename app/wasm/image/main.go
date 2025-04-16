//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"
)

func add(this js.Value, args []js.Value) interface{} {
	sum := args[0].Int() + args[1].Int()
	return js.ValueOf(sum)
}

// 写一个函数，获取浏览器的UserAgent
func getUserAgent(this js.Value, args []js.Value) interface{} {
	userAgent := js.Global().Get("navigator").Get("userAgent").String()
	setResult(userAgent)
	return js.ValueOf(userAgent)
}

func setResult(text string) {
	document := js.Global().Get("document")
	body := document.Call("getElementById", "result")
	body.Set("textContent", text)
}

func main() {
	c := make(chan struct{}, 0)

	// 将 Go 函数暴露给 JavaScript
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("getUserAgent", js.FuncOf(getUserAgent))

	// 获取 DOM 元素并设置文本内容
	setResult("Hello from Go!")

	<-c // 防止程序退出
}
