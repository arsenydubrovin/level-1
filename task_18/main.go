// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Inscrement() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.value
}

func main() {
	cnt := Counter{}

	// WaitGroup ждёт завершения всех горутин
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// Увеличение счётчика на 1 млн в каждой горутине
			for j := 0; j < 1_000_000; j++ {
				cnt.Inscrement()
			}
		}()
	}

	wg.Wait()

	// Вывод значения счётчика
	fmt.Println(cnt.Value())
}
