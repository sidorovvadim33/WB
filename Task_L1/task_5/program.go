package main

import (
	"context"
	"fmt"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

func main() {
	chanelInt := make(chan int)
	var seconds time.Duration

	fmt.Print("Введите количество секунд: ")
	scanln, err := fmt.Scanln(&seconds)
	//Проверка ввода на ошибки и кол-ва секунд больше 0
	if err != nil || scanln < 1 {
		return
	}

	/*
		Вариант с context.WithTimeout()

		Создаем контекст в который передаем время,
		канал Done возвращаемого контекста закрывается,
		когда вызывается возвращается функция cancel или когда
		закрывается канал Done родительского контекста,
		в зависимости от того, что произойдет раньше.
	*/
	ctx, cancel := context.WithTimeout(context.Background(), seconds*time.Second)
	defer cancel()

	go senderWithContext(ctx, chanelInt)

	for {
		foo, ok := <-chanelInt
		if !ok {
			break
		}
		println(foo)
	}
	fmt.Println("Время вышло")

	//Второй вариант с time.Tick
	//chanelInt = make(chan int)
	//stopCh := make(chan int)
	//
	////Время после которого нужно завершить выполнение: текущее + введенное в секундах
	//achievableTime := time.Now().Add(seconds * time.Second)
	//
	////Запускаем также передачу данных в канала через цикл
	//go sender(chanelInt, stopCh)
	//go func() {
	//	//Вывод всего что передано в канал
	//	for {
	//		fmt.Println(<-chanelInt)
	//	}
	//}()
	//
	//fmt.Println("Второй вариант")
	////Каждые 100 миллисекунд
	//for range time.Tick(100 * time.Millisecond) {
	//	currentTime := time.Now()
	//	difference := achievableTime.Sub(currentTime) //Вычисляем разницу между временем после которого программа завершается и текущим
	//
	//	total := int(difference.Seconds())
	//
	//	if total <= 0 { //Если разница меньше или равна нулю
	//		fmt.Println("Время вышло")
	//		stopCh <- 0 //Отправляем во второй канал 0 чтобы оператор select выполнил другой case
	//		os.Exit(0)
	//	}
	//}
}

//Передаем в качестве аргумента контекст и канал интов
func senderWithContext(ctx context.Context, ints chan int) {
	someNum := 0
	for {
		select {
		//Отправка в канал
		case ints <- someNum:
			someNum++
			//пока в канал Done не получить что-то на вход
		case <-ctx.Done():
			close(ints)
			return
		}
	}

}

func sender(ints, stop chan int) {
	result := 1
	for {
		select {
		case ints <- result:
			result++
		case <-stop:
			fmt.Println("Program stopped")
		}
	}
}
