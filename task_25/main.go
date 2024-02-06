// Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("программа стартовала")
	sleep(2 * time.Second)
	fmt.Println("прошло 2 секунды")
}

// sleep приостанавливает выполнении горутины на время d
func sleep(d time.Duration) {
	<-time.After(d)
}
