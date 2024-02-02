package points

import (
	"math"
)

// Point точка с координатами x и y.
type Point struct {
	x float64
	y float64
}

// New создаёт новый экземпляр Point.
func New(x, y float64) *Point {
	return &Point{x, y}
}

// CalculateDistance вычисляет расстояние между двумя точками.
func CalculateDistance(p1, p2 *Point) float64 {
	return math.Sqrt((p2.x-p1.x)*(p2.x-p1.x) + (p2.y-p1.y)*(p2.y-p1.y))
}
