/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

package main

import (
	"cmp"
	"fmt"
)

func main() {
	nums := []int{1, 3, 3, 10, 5, 1, 7, 12, 9, 10}
	quickSort(nums)
	fmt.Println(nums)

	strs := []string{"apple", "zuccini", "orange", "eggplant", "banana"}
	quickSort(strs)
	fmt.Println(strs)
}

func quickSort[T cmp.Ordered](arr []T) {
	if len(arr) <= 1 {
		return
	}

	if len(arr) == 2 {
		if arr[0] <= arr[1] {
			return
		}
		arr[0], arr[1] = arr[1], arr[0]
		return
	}

	left, right := 0, len(arr)-1
	pivot := arr[(left+right)/2]

	for {
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
