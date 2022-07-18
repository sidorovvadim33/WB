package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func main() {
	var x, y int
	fmt.Println("Введите x и y:")
	fmt.Scan(&x, &y)

	//Математический способ
	fmt.Printf("\nx:%d, y:%d\n", x, y)
	x = x + y
	y = x - y
	x = x - y
	fmt.Printf("x:%d, y:%d", x, y)

	//Способ с помощью особенностей языка
	x, y = y, x
	fmt.Printf("\nx:%d, y:%d\n", x, y)
}
