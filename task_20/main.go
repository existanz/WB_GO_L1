/*
Разработать программу, которая переворачивает слова в строке.

Пример: «snow dog sun — sun dog snow».
*/

package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(ReverseWords("now dog sun"))
}

func ReverseWords(sentence string) string {
	words := strings.Fields(sentence)
	slices.Reverse(words)
	return strings.Join(words, " ")
}
