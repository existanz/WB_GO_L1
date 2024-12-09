/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

package main

import (
	"fmt"
)

func main() {
	num := int64(185)
	fmt.Printf("%v \t --> \t %b\n", num, num)
	fmt.Printf("%v \t --> \t %b\n", setBit(num, 1), setBit(num, 1))
	fmt.Printf("%v \t --> \t %b\n", unsetBit(num, 7), unsetBit(num, 7))
	fmt.Printf("%v \t --> \t %b\n", unsetBit(num, 3), unsetBit(num, 3))
	fmt.Printf("%v \t --> \t %b\n", flipBit(num, 4), flipBit(num, 4))
	fmt.Printf("%v \t --> \t %b\n", flipBit(num, 200), flipBit(num, 200)) // выход за диапазон, но конструкция 1 << n при n > 63 вернёт 0, поскольку такого бита не существует
}

func setBit(num int64, i uint8) int64 {
	return num | (1 << i)
}

func unsetBit(num int64, i uint8) int64 {
	return num & ^(1 << i)
}

func flipBit(num int64, i uint8) int64 {
	return num ^ (1 << i)
}
