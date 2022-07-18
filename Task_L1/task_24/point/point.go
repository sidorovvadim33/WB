package point

import "math"

type Point struct {
	x int //x и y доступны только внутри пакета
	y int
}

func CreatePoint(x, y int) Point {
	//Возврщаем новый объект Point
	return Point{x, y}
}

func GetDistance(firstPoint, secondPoint Point) float64 {
	dx := secondPoint.x - firstPoint.x
	dy := secondPoint.y - firstPoint.y

	distance := math.Sqrt(float64(dx*dx + dy*dy))
	return distance
}
