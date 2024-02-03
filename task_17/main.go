// Реализовать бинарный поиск встроенными методами языка.

package main

import "fmt"

func main() {
	// слайс a отсортирован по возврастанию
	a := []int{1, 3, 8, 18, 24, 26, 28, 29, 30, 36, 37, 46, 49, 53, 59, 69, 72, 75, 77}
	fmt.Println(binarySearch(26, a))
}

// binarySearch ищет число n в слайсе a.
// Возвращает индекс и true, если число n есть в a, иначе возвращает ноль и false
func binarySearch(n int, a []int) (index int, exists bool) {
	l, r := 0, len(a)-1
	for l <= r {
		m := (l + r) / 2

		if a[m] == n {
			return m, true
		}

		if n > a[m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return 0, false
}
