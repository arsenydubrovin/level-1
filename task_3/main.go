// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов с использованием конкурентных вычислений.

package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}

	result := make(chan int) // Канал для передачи результатов вычисления горутины
	wg := sync.WaitGroup{}

	for _, n := range nums {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()

			result <- n * n // Запись квадрата числа в канал
		}(n)
	}

	// Запуск горутины для закрытия канала
	go func() {
		wg.Wait()
		close(result)
	}()

	// Вычисление суммы квадратов
	var sum int
	for n := range result {
		sum += n
	}

	fmt.Println(sum)
}
