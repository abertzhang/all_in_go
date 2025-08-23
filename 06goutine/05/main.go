package main
import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go worker(ctx)
	time.Sleep(10 * time.Second)
	fmt.Println("Main goroutine stopped")

}
func worker(ctx context.Context) {
	for{
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped")
			return
			default:
				fmt.Println("Worker is working...")
				time.Sleep(1 * time.Second)

		}
	}

}