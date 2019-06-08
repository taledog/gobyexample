package main

// 错误示例
var gvar int // 全局变量，声明不使用也可以

//func main() {
//	var one int     // error: one declared and not used
//	two := 2    // error: two declared and not used
//	var three int   // error: three declared and not used
//	three = 3
//}

// 正确示例
// 可以直接注释或移除未使用的变量
func main() {
	var one int
	_ = one

	two := 2
	println(two)

	var three int
	one = three

	var four int
	four = four
}
