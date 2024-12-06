/*
Реализовать конкурентную запись данных в map.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	const workers = 10

	m := make(map[int]string)
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	for id := range workers {
		wg.Add(1)
		go worker(wg, mu, id, m)
	}
	wg.Wait()

	fmt.Println(m)
}

func worker(wg *sync.WaitGroup, mu *sync.Mutex, id int, m map[int]string) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mu.Lock()
		m[id*10+i] = fmt.Sprintf("writen by worker %d", id)
		mu.Unlock()
	}
}
