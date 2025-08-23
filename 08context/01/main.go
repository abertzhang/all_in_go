package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go printNum(ctx, 5)

	time.Sleep(2 * time.Second)
	fmt.Println("主协程执行完毕,取消子协程")
}

func printNum(ctx context.Context,n int) {
	for i := 1; i < n; i++ {
		select {
			case <-ctx.Done():
			fmt.Println("协程被取消")
			default:
				fmt.Println("子协程中的数字",i)
				time.Sleep(500 * time.Millisecond)	
	}
}
}