package main

import (
	"fmt"
	"sort"
)

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	sort.Ints(items)
	fmt.Println(binarySearch(63, items))
}

// Логарифмическая сложность O(log n)
func binarySearch(item int, slice []int) bool {
	median := (len(slice)) / 2
	switch {
	// Если опорный элемент меньше чем item вызываем рекурсивно функцию со срезом от опорного и до конца
	case slice[median] < item:
		return binarySearch(item, slice[median+1:])
	case slice[median] > item:
		return binarySearch(item, slice[:median])
	case slice[median] == item:
		return true
	}
	return false
}
