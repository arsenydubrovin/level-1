// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

package main

import (
	"strings"
)

// createHugeString создаёт строку из n символов #
func createHugeString(n int) string {
	v := strings.Repeat("#", n)
	return v
}

func someFunc() string {
	v := createHugeString(1 << 10)

	// Срез от строки v[:100] оставляет неиспользованной большую часть выделенной памяти, но ссылается на неё, а значит сборщик мусора не будет её отчищать.

	// strings.Builder позволяет создать новую строку
	var b strings.Builder
	b.WriteString((v)[:100])

	// Использование глобальной переменной затрудняет отладку кода.
	// Лучше возвращать значение из функции
	return b.String()
}

func main() {
	justString := someFunc()

	justString = justString
}
