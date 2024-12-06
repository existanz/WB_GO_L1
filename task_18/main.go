/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	v  int
	mu sync.Mutex
}

func New() Counter {
	return Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.v++
	c.mu.Unlock()
}

func (c *Counter) Val() int {
	return c.v
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
