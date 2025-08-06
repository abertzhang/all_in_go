package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	deadlineCtx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
	operationDeadline(deadlineCtx)
	fmt.Println("Main: Exiting...")
}

func operationDeadline(ctx context.Context) {
	select {
	case <-ctx.Done():
		// 超时或取消
		fmt.Println("Operation deadline exceeded:", ctx.Err())
		return
	case <-time.After(3* time.Second):
		fmt.Println("Operation completed")
		return
	}
}