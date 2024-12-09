/*
Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
*/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "abcde"
	fmt.Println(str, "  \t", hasUniqueSymbols(str))

	str = "abcdeA"
	fmt.Println(str, "  \t", hasUniqueSymbols(str))

	str = "abCdefAaf"
	fmt.Println(str, "  \t", hasUniqueSymbols(str))

	str = "aabcde"
	fmt.Println(str, "  \t", hasUniqueSymbols(str))
}
func hasUniqueSymbols(s string) bool {
	hm := make(map[rune]struct{})
	for _, v := range s {
		v = unicode.ToLower(v) // регистронезависимость
		if _, ok := hm[v]; ok {
			return false
		}
		hm[v] = struct{}{}
	}
	return true
}
