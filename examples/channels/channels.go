// _通道(Channels)_ 是连接多个 Go 协程的管道。你可以从一个 Go 协程
// 将值发送到通道，然后在别的 Go 协程中接收。

package main

import "fmt"

func main() {

    // 使用 `make(chan val-type)` 创建一个新的通道。
    // 通道类型就是他们需要传递值的类型。
    messages := make(chan string)

    // 使用 `channel <-` 语法 _发送(send)_ 一个新的值到通道中。这里
    // 我们在一个新的 Go 协程中发送 `"ping"` 到上面创建的
    // `messages` 通道中。
    //for i := 0; i < 10; i++ {
    //    go func(value int) {
    //        print(value)
    //        messages <- "ping"
    //    }(i)
    //}

    //go func() {
        //for i := 0; i < 10; i++ {
            messages <- "ping"
        //}
    //}()

    // 使用 `<-channel` 语法从通道中 _接收(receives)_ 一个值。这里
    // 将接收我们在上面发送的 `"ping"` 消息并打印出来。

    //select {
    //case msg := <-messages:
    //    fmt.Printf(msg)
    //default:
    //   fmt.Printf("no communication\n")
    //}

    msg := <-messages
    fmt.Println(msg)
}
