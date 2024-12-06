/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 3, 10, 5, 1, 7, 12, 9, 10}
	quickSort(nums)
	fmt.Println(nums)
}

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	left, right := 0, len(arr)-1
	pivot := (arr[left] + arr[right]) / 2

	for left < right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left >= right {
			break
		}
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}

	quickSort(arr[:left])
	quickSort(arr[left:])
}
