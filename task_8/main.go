// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	var value int64 = 15

	fmt.Printf("%b\n", value)

	changedValue, err := setBit(value, 3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%b\n", changedValue)
}

func setBit(value int64, i int) (int64, error) {
	if i < 1 {
		return 0, errors.New("i should be >= 1")
	}

	if i > 63 {
		return 0, errors.New("i should be <= 63")
	}

	mask := int64(1 << (i - 1))

	return value ^ mask, nil
}
