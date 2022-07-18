package main

import "fmt"

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	var array [10]int
	for i := 0; i < len(array); i++ {
		array[i] = i
	}

	var firstChan = make(chan int)
	var secondChan = make(chan int)

	go func() {
		for _, elem := range array {
			firstChan <- elem
		}
	}()

	go channelProcessing(firstChan, secondChan)

	for i := 0; i < len(array); i++ {
		fmt.Println(<-secondChan)
	}
}

func channelProcessing(first, second chan int) {
	/*
		Функция принимает 2 канала, если первый канал открыт записывает увеличенное число во второй
	*/
	for {
		x, ok := <-first
		if !ok {
			break
		}
		second <- x * 2
	}
}
