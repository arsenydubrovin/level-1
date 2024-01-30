// Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

package main

import "fmt"

type human struct {
	id   int
	name string
}

type action struct {
	id    int
	human // Встраивание структуры
}

// Геттер для human.name
func (h *human) Name() string {
	return h.name
}

func main() {
	// Создание экземпляра структуры action
	a := action{
		id: 8,
		human: human{
			id:   5,
			name: "jeremy",
		},
	}

	fmt.Println(a.Name()) // Использование встроенного метода

	fmt.Println(a.id)       // Встроенное поле затеняется одноимённым
	fmt.Println(a.human.id) // Явное обращение к затенённой переменной
}
