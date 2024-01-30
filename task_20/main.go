// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func main() {
	words := "snow dog sun"
	fmt.Printf("%s — %s\n", words, reverseWords(words))
}

// С использованием string.Builder
func reverseWords(s string) string {
	words := strings.Fields(s)

	var b strings.Builder
	b.Grow(len(s))

	for len(words) > 0 {
		b.WriteString(" ")
		b.WriteString(words[len(words)-1]) // Извлечения последнего элемента из слайса
		words = words[:len(words)-1]       // Удаление последнего элемента в слайсе
	}

	return strings.TrimPrefix(b.String(), " ")
}

// С использованием "slices" и "strings"
// func reverseWords(s string) string {
// 	words := strings.Fields(s)
// 	slices.Reverse(words)

// 	return strings.Join(words, " ")
// }
