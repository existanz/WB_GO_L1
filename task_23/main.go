/*
Удалить i-ый элемент из слайса.
*/

package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(nums)
	nums = remove(nums, 5)
	fmt.Println(nums)
	nums = removeStable(nums, 5)
	fmt.Println(nums)
	nums = remove(nums, 10)
	fmt.Println(nums)
	nums = removeStable(nums, 10)
	fmt.Println(nums)

	strs := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	fmt.Println(strs)
	strs = remove(strs, 5)
	fmt.Println(strs)
	strs = removeStable(strs, 5)
	fmt.Println(strs)
	strs = remove(strs, 10)
	fmt.Println(strs)
	strs = removeStable(strs, 10)
	fmt.Println(strs)
}

// removeStable - удаляет i-й элемент из слайса, сохраняя порядок элементов в слайсе
// но он отработает за O(n)
func removeStable[T any](s []T, i int) []T {
	if i >= len(s) || i < 0 {
		return s
	}
	return append(s[:i], s[i+1:]...)
}

// remove - удаляет i-й элемент из слайса, не сохраняя порядок элементов в слайсе
// но он более эффективен и отработает за O(1)
func remove[T any](s []T, i int) []T {
	if i >= len(s) || i < 0 {
		return s
	}
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
