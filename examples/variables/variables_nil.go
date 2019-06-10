package main

import "fmt"

// 零值 panic
// Go 没有构造函数。正因为如此，它坚持认为“零值”应该是易于使用的。这是一个有趣的方法，但在我看来，它所带来的简化主要是针对语言实现者的。
// 还有一个很严重的问题: map 的零值:你可以查询它，但是在它里面存储东西会导致 panic:
func main() {
	fmt.Println("====================")
	fmt.Println()
	var m1 = map[string]string{} // 空值 map
	var m0 map[string]string     // 零值 map (nil)

	println(len(m1)) // 输出 '0'
	fmt.Println()
	fmt.Println("====================")
	println(len(m0))   // 输出 '0'
	println(m1["foo"]) // 输出 ''
	println(m0["foo"]) // 输出 ''
	m1["foo"] = "bar"  // 没问题
	//m0["foo"] = "bar"  // panics!
}
