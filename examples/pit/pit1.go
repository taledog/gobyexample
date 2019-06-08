package main

import "fmt"

//// 错误示例
//func main()
//{
//println("hello world")
//}
//
//// 等效于
//func main();    // 无函数体
//{
//println("hello world")
//}

// 正确示例
func main() {
	fmt.Println("hello world")
}

//// { 并不遵守分号注入规则，不会在其后边自动加分，此时可换行
//func main() {
//	{
//		println("hello world")
//	}
//}
