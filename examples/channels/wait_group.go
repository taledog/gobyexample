package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// 用 new 创建的对象返回的已经是指针
	//var wg = new(sync.WaitGroup)
	for i := 0; i < 5; i = i + 1 {
		wg.Add(1)
		go func(n int) {
			go EchoNumber(&wg, n)
		}(i)
	}
	wg.Wait()
}

func EchoNumber(wg *sync.WaitGroup, i int) {
	time.Sleep(3e9)
	fmt.Println(i)
	wg.Done()
}
