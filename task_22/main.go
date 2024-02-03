// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Два числа больших, чем 2^20
	a, b := big.NewInt(1<<34), big.NewInt(1<<24)

	fmt.Printf("%d + %d = %d\n", a, b, big.NewInt(0).Add(a, b))
	fmt.Printf("%d − %d = %d\n", a, b, big.NewInt(0).Sub(a, b))
	fmt.Printf("%d × %d = %d\n", a, b, big.NewInt(0).Mul(a, b))
	fmt.Printf("%d / %d = %d\n", a, b, big.NewInt(0).Div(a, b))
}
