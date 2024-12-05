/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

package main

import (
	"fmt"
)

func main() {
	var num int64
	fmt.Println("Input the number in range [-9223372036854775808, 9223372036854775807]")
	fmt.Scan(&num)

	var i uint8
	fmt.Println("Input the index in range [0, 62]")
	fmt.Scan(&i)

	if i == 63 {
		fmt.Println("You can't set the MSB")
	} else if i > 63 {
		fmt.Println("Index is out of range")
	}

	num ^= 1 << i
	fmt.Println(num)
}
