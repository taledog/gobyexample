// 在 Go 中，_变量_ 被显式声明，并可以被编译器用来
// 检查函数调用时的类型正确性。

package main

import "fmt"

func main() {

	// `var` 声明 1 个或者多个变量。
	var a = "initial"
	fmt.Println(a)

	// 你可以一次性声明多个变量。
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go 将自动推断已经初始化的变量类型。
	var d = true
	fmt.Println(d)

	// 声明后却没有给出对应的初始值时，变量将会初始化为
	// _零值_ 。例如，一个 `int` 的零值是 `0`。
	var e int
	fmt.Println(e)

	// `:=` 语法是声明并初始化变量的简写，例如
	// 这个例子中的 `var f string = "short"`。
	f := "short"
	fmt.Println(f)

	// 下面的代码会报错，因为对于引用类型的变量，我们不光要声明它，还要为它分配内容空间
	//var i *int
	//*i = 10
	//fmt.Println(*i)

	// 下面的代码不会报错，因为对于引用类型的变量，我们不光要声明它，因为为它分配了内容空间
	var j *int
	j = new(int)
	*j = 10
	fmt.Println(*j)

	// new和make就是为引用类型的变量分配内存空间的，但是对于本来就是引用类型的slice，chan和map来说，
	// 创建后它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，
	// 所以就没有必要返回他们的指针了，此时就用make来创建
}
