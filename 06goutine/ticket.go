package main

import (
	"fmt"
	"sync"
)

func main() {
	tickets := 500
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(userId int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			if tickets > 0 {
				tickets--
				fmt.Printf("用户抢到票,剩余票数:%d\n", tickets)
			} else {
				fmt.Printf("票已卖完\n")
			}
		}(i)
	}
	wg.Wait()
}
