// Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

func main() {
	// два неупорядоченных множества
	m := []int{3, 44, 84, 1, 0, 42, -3}
	n := []int{1, 74, 0, 42, 3, 2}

	var intersection []int

	// мапа для быстрой проверки
	set := make(map[int]struct{})

	for _, el := range m {
		set[el] = struct{}{}
	}

	for _, el := range n {
		if _, exists := set[el]; exists {
			intersection = append(intersection, el)
		}
	}

	fmt.Println(intersection)
}
