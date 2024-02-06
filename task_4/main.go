// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

const (
	workersAmount = 10 // Настройка количества воркеров
)

func main() {
	var wg sync.WaitGroup

	// Канал для передачи произвольных данных
	data := make(chan any)

	// Канал для передачи сигнала о завершении
	stop := make(chan struct{})

	for i := 1; i <= workersAmount; i++ {
		wg.Add(1)
		// Запуск воркера (WaitGroup передаётся по указателю)
		go runWorker(i, data, stop, &wg)
	}

	// Запуск горутины, пишущей в канал
	go runWriter(data, stop)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	// Ожидание перывания программы через ctrl+C
	<-sig
	close(stop) // Все горутины получают сигнал о завершении
	wg.Wait()
}

// runWriter запускает цикл, который каждую секунду пишет данные в канал data
func runWriter(data chan any, stop chan struct{}) {
	defer close(data) // Пишущая горутина ответственная за закрытие канала

	for {
		select {
		case <-stop:
			return

		case data <- time.Now().String(): // Каждую секунду в канал отправляются данные
			time.Sleep(1 * time.Second)
		}
	}
}

// runWorker запускает воркер, который ждёт сообщений из канала data и выводит их в stdout
func runWorker(id int, data chan any, stop chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-stop:
			log.Printf("воркер %d\tполучен сигнал о завершении\n", id)
			return

		case msg, ok := <-data:
			if !ok {
				log.Printf("воркер %d\tканал закрыт\n", id)
				return
			}

			fmt.Printf("воркер %d\tполучено сообщение: %s\n", id, msg)
		}
	}
}
