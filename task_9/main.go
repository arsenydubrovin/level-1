// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
	"sync"
)

// writeFromArray отправляет числа из nums в канал out
func writeFromArray(nums []int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out)

	for _, num := range nums {
		out <- num
	}
}

// writeFromChannel читает числа из канала in, умножает их на два и пишет в канал out
func writeFromChannel(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out)

	for num := range in {
		out <- num * 2
	}
}

// readFromChannel читает из канала in и пишет stdout
func readFromChannel(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for n := range in {
		fmt.Println(n)
	}
}

func main() {
	// Создание каналов для передачи данных между горутинами
	channel1 := make(chan int)
	channel2 := make(chan int)

	// WaitGroup для ожидания трёх горутин
	var wg sync.WaitGroup
	wg.Add(3)

	arr := []int{10, 7, 8, 11, 15, 1, 2, 14, 6, 3, 9, 13, 12, 5, 4}

	// Запуск  конвеера
	go writeFromArray(arr, channel1, &wg)
	go writeFromChannel(channel1, channel2, &wg)
	go readFromChannel(channel2, &wg)

	wg.Wait()
}
