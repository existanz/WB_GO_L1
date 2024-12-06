/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// проще было бы `type Counter atomic.Uint64`
// но тогда решение не соответствовало условию "реализовать структуру-счетчик"
type Counter struct {
	av atomic.Uint64
}

func New() Counter {
	return Counter{}
}

func (c *Counter) Inc() {
	c.av.Add(1)
}

func (c *Counter) Val() uint64 {
	return c.av.Load()
}

func main() {
	wg := sync.WaitGroup{}
	c := Counter{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg, &c)
	}
	wg.Wait()
	fmt.Println("Counter value is:", c.Val())
}

func worker(wg *sync.WaitGroup, c *Counter) {
	defer wg.Done()
	for range 10 {
		c.Inc()
	}
}
