package main

import (
	"context"
	"fmt"
	"time"
)
func main() {
timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()
operationTimeout(timeoutCtx)
fmt.Println("Main: Exiting...")
}
func operationTimeout(ctx context.Context) {
select {
	case <-ctx.Done():
		// 超时或取消
		fmt.Println("Operation timed out:", ctx.Err())
		return
	case <-time.After(10 * time.Second):
		fmt.Println("Operation completed")
		return
}
}