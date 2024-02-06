// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
)

const (
	goroutinesAmount = 20
)

// 1. С использованием стандартной мапы и мьютекса
// func main() {
// 	var wg sync.WaitGroup
// 	var mu sync.Mutex

// 	data := make(map[int]struct{})

// 	wg.Add(goroutinesAmount) // Увеличение счётчика WaitGroup
// 	for i := 0; i < goroutinesAmount; i++ {
// 		go func(i int) {
// 			defer wg.Done()

// 			// Блокировка перед записью в мапу
// 			mu.Lock()
// 			defer mu.Unlock() // Разблокировка после завершения горутины

// 			data[i] = struct{}{}
// 		}(i)
// 	}

// 	wg.Wait()

// 	// Вывод всех ключей в мапе
// 	for key := range data {
// 		fmt.Printf("%d ", key)
// 	}
// 	fmt.Println()
// }

// 2. С использованием sync.Map
func main() {
	var wg sync.WaitGroup

	var data sync.Map

	wg.Add(goroutinesAmount) // Увеличение счётчика WaitGroup
	for i := 0; i < goroutinesAmount; i++ {
		go func(i int) {
			defer wg.Done()

			data.Store(i, struct{}{})
		}(i)
	}

	wg.Wait()

	// Вывод всех ключей в мапе
	data.Range(func(key, value any) bool {
		fmt.Printf("%d ", key)
		return true
	})
	fmt.Println()
}
