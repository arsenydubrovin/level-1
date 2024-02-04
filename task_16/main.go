// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import "fmt"

func main() {
	arr := []int{189, 72, 33, 63, 11, 174, -60, -138}
	fmt.Printf("Неотсортированный слайс: %v\n", arr)

	quickSort(arr, 0, len(arr)-1)
	fmt.Printf("Отсортированный слайс: %v\n", arr)
}

func quickSort(arr []int, l, r int) {
	if l < r {
		pivotIndex := partition(arr, l, r)

		quickSort(arr, l, pivotIndex-1)
		quickSort(arr, pivotIndex+1, r)
	}
}

func partition(arr []int, l, r int) int {
	pivot := arr[r]
	i := (l - 1)

	for j := l; j <= r-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[r] = arr[r], arr[i+1]

	return (i + 1)
}
