/*
Реализовать все возможные способы остановки выполнения горутины.
*/

package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
	workingTime   = 3 * time.Second
	operationTime = 500 * time.Millisecond
)

func main() {
	wg := &sync.WaitGroup{}

	// самый простой способ использование флага
	// 1.1 Булевый

	flag := false
	wg.Add(1)
	go workerWithFlag(wg, &flag)

	time.Sleep(workingTime)
	flag = true

	wg.Wait()

	// при таком подходе есть риск рассинхрона
	// чтобы этого избежать можно использовать примитивы синхронизации
	// 1.2 Mutex

	var f = Flag{flag: new(bool), mu: new(sync.Mutex)}

	wg.Add(1)
	go workerWithStruct(wg, f)

	time.Sleep(workingTime)

	f.mu.Lock()
	*f.flag = true
	f.mu.Unlock()

	wg.Wait()

	// 1.3 Atomic

	var atomicFlag = new(atomic.Bool) // )
	atomicFlag.Store(false)

	wg.Add(1)
	go workerWithAtomic(wg, atomicFlag)

	time.Sleep(workingTime)

	atomicFlag.Store(true)

	wg.Wait()

	// 2. Классический вариант с done-каналом

	done := make(chan bool)

	wg.Add(1)
	go workerWithDone(wg, done)

	time.Sleep(workingTime)

	done <- true

	wg.Wait()

	// 3. Контекст

	// 3.1 WithCancel

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go workerWithContext(ctx, wg)

	time.Sleep(workingTime)

	cancel()

	wg.Wait()

	// 3.2 WithTimeout
	ctx, cancel = context.WithTimeout(context.Background(), workingTime)

	wg.Add(1)
	go workerWithContext(ctx, wg)

	wg.Wait()

	// можно ещё применить WithDeadline, но под капотом это то же самое что и WithTimeout
	// или сигналы signal.NotifyContext, что то же самое что и WithCancel

	fmt.Println("Stopped main goroutine")

}

type Flag struct {
	flag *bool
	mu   *sync.Mutex
}

func workerWithFlag(wg *sync.WaitGroup, flag *bool) {
	defer wg.Done()
	fmt.Println("Worker with flag started")
	for !*flag {
		fmt.Println("Worker with flag do some work...")
		time.Sleep(operationTime)
	}
	fmt.Println("Worker with flag stopped")
}

func workerWithStruct(wg *sync.WaitGroup, f Flag) {
	defer wg.Done()
	fmt.Println("Worker with mutex started")
	for {
		f.mu.Lock()
		flag := *f.flag
		f.mu.Unlock()
		if flag {
			break
		}
		fmt.Println("Worker with mutex do some work...")
		time.Sleep(operationTime)
	}
	fmt.Println("Worker with mutex stopped")
}

func workerWithAtomic(wg *sync.WaitGroup, flag *atomic.Bool) {
	defer wg.Done()
	fmt.Println("Worker with atomic started")
	for !flag.Load() {
		fmt.Println("Worker with atomic do some work...")
		time.Sleep(operationTime)
	}
	fmt.Println("Worker with atomic stopped")
}

func workerWithDone(wg *sync.WaitGroup, done <-chan bool) {
	defer wg.Done()
	fmt.Println("Worker with done started")
	for {
		select {
		case <-done:
			fmt.Println("Worker with done stopped")
			return
		default:
			fmt.Println("Worker with done do some work...")
			time.Sleep(operationTime)
		}
	}
}

func workerWithContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker with context started")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker with context stopped")
			return
		default:
			fmt.Println("Worker with context do some work...")
			time.Sleep(operationTime)
		}
	}
}
