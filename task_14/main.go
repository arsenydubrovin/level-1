// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// i может иметь тип int, string, bool или channel
	var i interface{} = make(<-chan int)

	// type switch с использованием рефлексии
	switch reflect.TypeOf(i).Kind() {
	case reflect.Int:
		fmt.Println("int")
	case reflect.String:
		fmt.Println("string")
	case reflect.Bool:
		fmt.Println("bool")
	case reflect.Chan:
		fmt.Println("chan")
	default:
		fmt.Println("неизвестный тип")
	}
}
