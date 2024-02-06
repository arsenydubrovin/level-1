// Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

// 1. Сигнал через канал stop
func main() {
	var wg sync.WaitGroup

	stop := make(chan struct{})

	t := time.NewTicker(500 * time.Millisecond)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-stop:
				// Завершение работы
				fmt.Println("получен сигнал stop")
				return

			case <-t.C:
				// Имитация работы горутины
				fmt.Println("горутина работает")
			}
		}
	}()

	// Ожидание перывания программы через ctrl+C
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	// Закрытие канала stop для передачи сигнала о завершении горутине
	close(stop)
	wg.Wait()
}
