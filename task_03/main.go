/*
Дана последовательность чисел: 2,4,6,8,10.

Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	nums := []int64{2, 4, 6, 8, 10}
	var sum atomic.Int64 // по идее мы могли использовать mutex вместо atomic, но так будет быстрее

	wg := &sync.WaitGroup{}

	for _, num := range nums {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sqr := num * num
			sum.Add(sqr)
		}()
	}
	wg.Wait()
	fmt.Println(sum.Load())
}
