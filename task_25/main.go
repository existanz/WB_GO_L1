/*
Реализовать собственную функцию sleep.
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Start", time.Now())

	MySleep(5 * time.Second)

	fmt.Println("End", time.Now())
}

func MySleep(d time.Duration) {
	// реализовать sleep без библиотеки time будет довольно сложно
	// поэтому используем её, тем более, что по условию она не запрещена
	// самый эффективный вариант, помимо непосредственного использования
	// time.Sleep(d) будет использование канала time.After(d)
	<-time.After(d)
}
