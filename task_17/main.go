/*
Реализовать бинарный поиск встроенными методами языка.
*/

package main

import (
	"cmp"
	"fmt"
)

func main() {
	nums := []int{1, 2, 13, 14, 25, 26, 47, 81, 98}
	fmt.Println(BinarySearch(nums, 15))
	fmt.Println(BinarySearch(nums, 26))

	strs := []string{"apple", "banana", "carrot", "durian", "eggplant", "fig", "guarana", "hazelnut"}
	fmt.Println(BinarySearch(strs, "dragonfruit"))
	fmt.Println(BinarySearch(strs, "fig"))
}

// Бинарный поиск возвращает индекс, по которому в слайсе находится переданное значение,
// либо если значение не найдено, по которому можно это значение вставить.
func BinarySearch[T cmp.Ordered](arr []T, x T) int {
	n := len(arr)

	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if arr[mid] < x {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
