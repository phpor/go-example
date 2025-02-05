package main

import (
	"bytes"
	"fmt"
)

// 使用说明： 给定一个以字符串表示的非负整数  num，移除这个数中的 k 位数字，使得剩下的数字最大。
// https://leetcode.cn/problems/remove-k-digits/solution/yi-zhao-chi-bian-li-kou-si-dao-ti-ma-ma-zai-ye-b-5/

func removeKDigits(num string, k int) string {
	var stack []byte

	for i := 0; i < len(num); i++ {
		digit := num[i]
		// 如果栈顶元素小于当前元素，并且可以删除栈顶元素
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] < digit {
			stack = stack[:len(stack)-1] // 弹出栈顶元素
			k--
		}
		stack = append(stack, digit)
		println(string(stack))
	}

	// 如果还有剩余的删除次数，继续从栈中弹出元素
	for k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}

	// 构建结果字符串，并去掉前导零
	result := bytes.TrimLeft(stack, "0")
	if len(result) == 0 {
		return "0"
	}
	return string(result)
}

// 仅仅前后比较是不够的，需要回溯，必须使用栈
func getMax(num string, k int) string {
	var result = make([]byte, len(num)-k)
	j := 0
	for i := 0; i < len(num); i++ {
		if k == 0 {
			result[j] = num[i]
			j++
			continue
		}
		if i < len(num)-1 && num[i] < num[i+1] {
			k--
			continue
		}
		result[j] = num[i]
		j++
	}
	return string(result)
}

func main() {
	num := "27886483"
	m := 4

	result := removeKDigits(num, m)
	//result2 := getMax(num, m)
	fmt.Printf("删除 %d 个数字后，最大的数是: %s\n", m, result)
	//fmt.Printf("删除 %d 个数字后，最大的数是: %s\n", m, result2)
}
