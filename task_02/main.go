package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := [...]int{2, 4, 6, 8, 10}

	sqrs := make(chan int) // создаём небуферизированный канал

	wg := &sync.WaitGroup{} // используем wg для синхронизации

	for _, num := range nums {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sqrs <- num * num
		}()
	}
	go func() {
		wg.Wait()
		close(sqrs)
	}()

	for sqr := range sqrs { // range по каналу идёт пока канал не закрыт
		fmt.Println(sqr)
	}

}
