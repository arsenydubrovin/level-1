// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

type Set struct {
	data map[string]struct{}
}

// NewSet создаёт новый экземпляр Set
func NewSet() *Set {
	return &Set{
		data: make(map[string]struct{}),
	}
}

// Add добавляет значение в множество
func (s *Set) Add(value string) {
	s.data[value] = struct{}{}
}

// Pop возвращает true, если значение есть в множестве, и удаляет его из множества
func (s *Set) Pop(value string) bool {
	_, exists := s.data[value]
	delete(s.data, value)

	return exists
}

// List возвращает список всех значений в множестве
func (s *Set) List() []string {
	list := make([]string, 0, len(s.data))

	for value := range s.data {
		list = append(list, value)
	}

	return list
}

func main() {
	set := NewSet()

	// Добавление элементов
	set.Add("cat")
	set.Add("cat")
	set.Add("dog")
	set.Add("cat")
	set.Add("tree")

	fmt.Println(set.List())

	// Удаление элементов
	set.Pop("cat")
	set.Pop("dog")

	fmt.Println(set.List())
}
