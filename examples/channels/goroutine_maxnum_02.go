package main

import (
	"fmt"
	"strconv"
	"time"
)

var (
	maxRoutineNum = 10
)

// 模拟下载页面的方法
func download(url string, ch chan int) {
	fmt.Println("download from ", url)
	// 休眠两秒模拟下载页面
	time.Sleep(2 * 1e9)
	// 下载完成则从ch推出数据
	<-ch
}

func main() {
	ch := make(chan int, maxRoutineNum)

	urls := [100]string{}
	for i := 0; i < 100; i++ {
		urls[i] = "url" + strconv.Itoa(i)
	}
	for i := 0; i < len(urls); i++ {
		// 开启下载协程前往ch塞一个数据
		// 如果ch满了则会处于阻塞，从而达到限制最大协程的功能
		ch <- 1
		go download(urls[i], ch)
	}

	// 休眠一下
	for {
		time.Sleep(1 * 1e9)
	}
}
