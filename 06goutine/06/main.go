package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	worker := NewAdvanceWorker()
	worker.Start()
	time.Sleep(5 * time.Second)
	worker.Stop()
	fmt.Println("Main goroutine stopped")

}

type AdvanceWorker struct {
	ctx      context.Context
	ctxTimeout context.Context
	cancel   context.CancelFunc
	wg       sync.WaitGroup
	stopChan chan struct{}
}

func NewAdvanceWorker() *AdvanceWorker {
	ctx, cancel := context.WithCancel(context.Background())
	timeoutCtx, _ := context.WithTimeout(ctx, 10*time.Second)
	return &AdvanceWorker{
		ctx:      ctx,
		ctxTimeout: timeoutCtx,
		cancel:   cancel,
		stopChan: make(chan struct{}),
	}
}
func (w *AdvanceWorker) Start() {
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-w.ctx.Done():
				fmt.Println("Worker stopped by ctx")
				return
			case <-w.stopChan:
				fmt.Println("Worker stopped by stopChan")
			case <-ticker.C:
				fmt.Println("Worker is working...")
			}
		}
	}()
}
func (w *AdvanceWorker) Stop() {
	w.cancel()
	close(w.stopChan)
	w.wg.Wait()
}
