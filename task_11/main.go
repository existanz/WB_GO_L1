/*
Реализовать пересечение двух неупорядоченных множеств.
*/

package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{2, 3, 4, 5, 6}

	fmt.Println(merge(a, b))

	c := []string{"apple", "banana", "durian", "eggplant", "cocoa"}
	d := []string{"zuccini", "cocoa", "durian", "fig", "eggplant"}
	fmt.Println(merge(c, d))
}

func merge[T comparable](a, b []T) []T {
	hm := make(map[T]int)
	out := make([]T, 0, min(len(a), len(b)))
	for _, v := range a {
		hm[v]++
	}

	for _, v := range b {
		if hm[v] > 0 {
			out = append(out, v)
			hm[v]--
		}
	}

	return out
}
