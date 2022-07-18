package main

import (
	"fmt"
	"sort"
)

type Slice struct {
	slice []int
}

func (s *Slice) Len() int {
	return len(s.slice)
}

func (s *Slice) Less(i, j int) bool {
	return s.slice[i] < s.slice[j]
}

func (s *Slice) Swap(i, j int) {
	s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
}

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

сложность O(n*log(n))
*/

func main() {
	var slice = []int{5, 1, 0, 7, 4, 3, 10, 9, -1, -5}
	fmt.Println("После реализации алгоритма:", quickSort(slice))

	slice = []int{5, 1, 0, 7, 4, 3, 10, 9, -1, -5}
	fmt.Println("До метода sort.Sort():", slice)
	//Функция сорт принимает на вход interface у которого необходимо реализовать 3 метода
	sort.Sort(&Slice{slice})
	fmt.Println("После:", slice)

	slice = []int{5, 1, 0, 7, 4, 3, 10, 9, -1, -5}
	fmt.Println("До метода sort.Ints():", slice)
	//Нет необходимости реализации методов
	sort.Ints(slice)
	fmt.Println("После:", slice)
}

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := len(a) / 2

	// Перемещаем опорное значение в конец среза
	a[pivot], a[right] = a[right], a[pivot]

	//Перебор среза
	for i, _ := range a {
		//Если текущее значение меньше опорного, которое стоит в конце
		if a[i] < a[right] {
			//То мы меняем местами левую границу и текущий элемент
			a[left], a[i] = a[i], a[left]
			//Увеличиваем значение левой границы
			left++
		}
	}
	/*
		В данном цикле по сути за счет left мы перекидываем в левую часть массива все элементы меньше опорного,
		потому что left у нас увеличивается только при перемещении, а все элементы которые встречались в левой части,
		но больше опорного окажутся где-то на месте a[i], которое точно дальше чем left
	*/

	//Перемещаем значение опорного элемента на место левой границы, получаем срез в котором элементы слева меньше опорного, а справа больше
	a[left], a[right] = a[right], a[left]

	// Рекурсивный вызов сортировки от начала до опорного и т.д.
	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}
