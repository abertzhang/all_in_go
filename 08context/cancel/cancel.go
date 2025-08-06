package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	parentCtx, cancel := context.WithCancel(context.Background())
	go worker(parentCtx)
	time.Sleep(2 * time.Second)
	cancel()
	// time.Sleep(1 * time.Second)
	fmt.Println("Main: Exiting...")
}
func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker: Received cancellation signal, exiting...")
			return
		default:
			fmt.Println("Worker: Working...")
			time.Sleep(500 * time.Millisecond) // 模拟工作
		}
	}
}