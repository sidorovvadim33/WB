package main

import (
	point "Task_L1/task_24/point"
	"fmt"
)

/*
Разработать программу нахождения расстояния между двумя точками, которые
представлены в виде структуры Point с инкапсулированными параметрами x,y
и конструктором.
*/

func main() {
	p1 := point.CreatePoint(2, 4)
	p2 := point.CreatePoint(25, -54)

	fmt.Println(point.GetDistance(p1, p2))
}
