/*
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	nums := [...]int{1, 2, 3, 4, 5}

	go func() {
		for _, num := range nums {
			ch1 <- num
		}
		close(ch1)
	}()

	go func() {
		for num := range ch1 {
			ch2 <- num * 2
		}
		close(ch2)
	}()

	for num := range ch2 {
		fmt.Println(num)
		time.Sleep(time.Second)
	}
}
