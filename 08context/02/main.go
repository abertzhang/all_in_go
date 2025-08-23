package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	go printNum(stopCh,10)
	time.Sleep(3 * time.Second)
	close(stopCh)
	time.Sleep(2 * time.Second)
	fmt.Println("主协程执行完毕")
}

func printNum(stopCh chan struct{},n int) {
	for i := 1; i <= n; i++ {
		select {
		case <-stopCh:
			fmt.Println("协程被取消")
			return
		default:
			fmt.Println("子协程中的数字", i)
			time.Sleep(500 * time.Millisecond)
		}
	}
}