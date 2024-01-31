// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

package main

import (
	"fmt"
	"strings"
)

func main() {
	testStrings := []string{"abcd", "abCdefA", "aabcd"}

	for _, str := range testStrings {
		fmt.Println(str, checkUnique(str))
	}
}

// checkUnique проверяет уникальность символов в строке вне зависимости от регистра
func checkUnique(s string) bool {
	s = strings.ToLower(s)

	// Множество символов, которые встречались в строке
	chars := make(map[rune]struct{})

	for _, c := range s {
		if _, met := chars[c]; met {
			return false
		}
		chars[c] = struct{}{}
	}

	return true
}
