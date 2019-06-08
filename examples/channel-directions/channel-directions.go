// 当使用通道作为函数的参数时，你可以指定这个通道是不是
// 只用来发送或者接收值。这个特性提升了程序的类型安全性。

package main

import "fmt"

// `ping` 函数定义了一个只允许发送数据的通道。尝试使用这个通
// 道来接收数据将会得到一个编译时错误。
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// `pong` 函数允许通道（`pings`）来接收数据，另一通道
// （`pongs`）来发送数据。
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	// 通道被定义成只写，也就是只能向通道发送数据
	var a chan<- string
	a = make(chan<- string, 1)
	println(a)

	// 通道被定义成只读，也就是只能从通道里读数据
	var b <-chan string
	b = make(<-chan string, 1)
	println(b)

	// 通道被定义成读写，也就是通道可以读也可以写
	var c chan string
	c = make(chan string, 1)
	println(c)

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
