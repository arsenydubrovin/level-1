// Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"

	"github.com/arsenydubrovin/level-1/task_24/points"
)

func main() {
	pointA := points.New(3.2, 5.3)
	pointB := points.New(2.623, -1.125)

	fmt.Println(points.CalculateDistance(pointA, pointB))
}
