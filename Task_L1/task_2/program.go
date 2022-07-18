package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает
значение квадратов чисел взятых из массива
(2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	var arr = [5]int{2, 4, 6, 8, 10}

	fmt.Println("Первый способ")
	//Определяем группу горутин
	var wg sync.WaitGroup
	wg.Add(len(arr))

	//Перебор массива через range
	for _, elem := range arr {
		// Создаем горутину для каждого числа массива
		go squareWithWG(elem, &wg)
	}
	//Ожидаем завершения группы
	wg.Wait()

	fmt.Println("\nВторой способ:")
	//Инициализируем канал
	chanelInt := make(chan int, len(arr))
	// С помощью оператора defer в конце main закрываем канал
	defer close(chanelInt)

	//Перебор массива через range
	for _, elem := range arr {
		// Создаем горутину для каждого числа массива
		go squareWithChannel(elem, chanelInt)
	}

	for range arr {
		//Вывод переданных в канал значений
		fmt.Println(<-chanelInt)
	}
}

func squareWithWG(number int, wg *sync.WaitGroup) {
	//Вывод квадрата переданного числа
	fmt.Println(number * number)
	//Уменьшаем счетчик горутин группы
	wg.Done()
}

func squareWithChannel(number int, chanelInt chan int) {
	//Передаем в канал квадрат числа
	chanelInt <- number * number
}
