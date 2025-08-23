package main

import (
	"fmt"
	"time"
)

func main() {
	stopChan := make(chan bool)
	go worker(stopChan)
	time.Sleep(5 * time.Second)
	stopChan <- true
	time.Sleep(1 * time.Second)
	fmt.Println("Main goroutine stopped")
}
func worker(stopChan chan bool) {
	for {
		select {
		case <-stopChan:
			fmt.Println("Worker stopped")
			return
		default:
			fmt.Println("Worker is working...")
			time.Sleep(1 * time.Second)
		}
	}
}
