package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i = i + 1 {
		wg.Add(1)
		go func(n int) {
			// defer wg.Done()
			defer wg.Add(-1)
			go EchoNumber(wg, n)
		}(i)
	}
	wg.Wait()
}

func EchoNumber(wg sync.WaitGroup, i int) {
	time.Sleep(3e9)
	fmt.Println(i)
}
