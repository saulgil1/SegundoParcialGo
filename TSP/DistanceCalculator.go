package TSP

import (
	"math"
)

type Point struct {
	X float64
	Y float64
}

func EuclideanDistance(p1 Point, p2 Point) int {
	distancia := int(math.Sqrt((((p2.X - p1.X) * (p2.X - p1.X)) + ((p2.Y - p1.Y) * (p2.Y - p1.Y)))))
	return distancia
}
