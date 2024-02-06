// Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

package main

import (
	"fmt"
	"time"
)

const (
	timeout = 5 // секунд
)

func main() {
	ch := make(chan int)

	// Запуск горутин для чтения и записи
	go read(ch)
	go send(ch)

	// Завершение по таймеру

	// 1. При помощи Sleep
	// 	time.Sleep(timeout * time.Second)

	// 2. При помощи time.Timer
	// timer := time.NewTimer(timeout * time.Second)
	// <-timer.C

	// 3. При помощи time.After
	<-time.After(timeout * time.Second)

	fmt.Printf("прошло %d секунд. завершение программы\n", timeout)
}

// send каждые полсекунды отправляет значение в ch
func send(ch chan<- int) {
	defer close(ch)
	for {
		ch <- int(time.Now().UnixMilli())
		time.Sleep(500 * time.Millisecond)
	}
}

// read чистает значения из ch
func read(ch <-chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("канал закрыт")
			return
		}

		fmt.Printf("получено значение: %d\n", v)
	}
}
