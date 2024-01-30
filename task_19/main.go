// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import (
	"fmt"
	// "slices"
)

func main() {
	word := "главрыба"
	fmt.Printf("%s — %s\n", word, reverseString(word))
}

// Без использование пакетов
func reverseString(s string) string {
	runes := []rune(s)

	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}

	return string(runes)
}

// С использованием slices
// func reverseString(s string) string {
// 	runes := []rune(s)
// 	slices.Reverse(runes)

// 	return string(runes)
// }
