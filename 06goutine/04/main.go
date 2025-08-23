package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var stopFlag int32 = 0
	go worker(&stopFlag)
	time.Sleep(5 * time.Second)
	atomic.StoreInt32(&stopFlag, 1)
	time.Sleep(1 * time.Second)
	fmt.Println("Main goroutine stopped")
}
func worker(stopFlag *int32) {
	for  {
		if(atomic.LoadInt32(stopFlag) == 1){
			fmt.Println("Worker stopped")
			return
		}
		fmt.Println("Worker is working...")
		time.Sleep(1 * time.Second)
	}

}