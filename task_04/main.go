/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/

package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var workers int
	fmt.Println("Input the number of workers")
	fmt.Scan(&workers)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	wg := &sync.WaitGroup{}

	for i := range workers {
		wg.Add(1)
		go worker(ctx, wg, i+1, t.C)
	}

	<-ctx.Done()
	wg.Wait()
}

func worker(ctx context.Context, wg *sync.WaitGroup, workerID int, ch <-chan time.Time) {
	defer wg.Done()

	fmt.Printf("Worker %d started\n", workerID)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopped by context\n", workerID)
			return
		case data := <-ch:
			fmt.Printf("Worker %d received data: %v\n", workerID, data)
		}
	}
}
