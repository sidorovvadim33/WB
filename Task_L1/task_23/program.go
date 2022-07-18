package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var num int
	fmt.Scan(&num)
	fmt.Println(removeElem(slice, num))
}

/*
Просто соединяем два среза: от начала и до элемента не включительно и от следующего от элемента до конца.
*/
func removeElem(slice []int, num int) []int {
	return append(slice[:num], slice[num+1:]...)
}
