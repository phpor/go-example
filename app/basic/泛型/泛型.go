package main

// 泛型
// 对于Stringify函数，泛型的收益在于静态类型检查，至于执行性能应该没有改善
// 对于Max函数，泛型的使用使得不需要做类型转换，执行性能应该有所改善
func main() {
	println(Stringify(1))
	println(Stringify("one"))
	println(Max([]int{1, 2, 3}))
	println(Max([]string{"a", "b", "c"}))
}

// Stringify 使用泛型写一个函数，接收字符串或整型，返回字符串，如果是字符串，返回字符串本身，如果是整型，返回整型的字符串形式
func Stringify[T int | string](s T) string {
	switch v := any(s).(type) {
	case string:
		return v
	case int:
		return string(v)
	default:
		return ""
	}
}

// Max 使用泛型写一个函数，接收一个int或string的slice，返回这个slice中的最大的那个元素
func Max[T int | string](s []T) T {
	var max T
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max
}
