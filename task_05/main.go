/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	const N = 5
	ch := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), N*time.Second)
	defer cancel()

	go func() {
		defer close(ch)
		i := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Timeout")
				return
			case ch <- i:
				i++
			}
			time.Sleep(time.Millisecond * 300) // чтобы не печаталось слишком быстро
		}
	}()

	for v := range ch { // range по каналу проходит пока канал не закрыт
		fmt.Println(v)
	}
}
