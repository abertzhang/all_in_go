package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	stopChan := make(chan struct{})
	wg.Add(1)
	go worker(&wg, stopChan)
	time.Sleep(10 * time.Second)
	close(stopChan)
	wg.Wait()
	fmt.Println("Main goroutine stopped")
}

func worker(wg *sync.WaitGroup, stopChan chan struct{}) {
	defer wg.Done()
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
