package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

/*
6. Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {
	/*
		1 вариант:
		Остановка выполнения горутины, после того как во второй
		канал передано знаечение, проверка через оператор select
	*/
	quit := make(chan bool)
	output := make(chan string)

	go selectStop(output, quit)

	for i := 0; i < 20; i++ {
		fmt.Println(<-output)
	}
	quit <- true
	fmt.Println("Горутина 1 — всё")
	close(quit)
	close(output)

	//2 вариант
	/*
		2 вариант
		Проверка того что канал еще открыл, если нет: завершаем горутину

		В цикле передаем 20 значение после чего вызов close(intChanel)
	*/
	wg := sync.WaitGroup{}
	wg.Add(1)
	intChanel := make(chan int)

	go chanCheckStop(intChanel, &wg)

	for i := 0; i < 20; i++ {
		intChanel <- i
	}
	//Закрыли канал — внутри горутины проверка не пройдена, завершаем ее через wg.Done()
	close(intChanel)

	wg.Wait()
	fmt.Println("Горутина 2 — всё")

	/*
		3 вариант

		Аналогичен первому и использует оператор select при вызове cancel()
		в канал Done передаются данные из-за чего выполнение горутины завершается
	*/
	output = make(chan string)
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // if cancel() execute
				output <- "Горутина 3 — работаем"
				return
			}
		}
	}(ctx)

	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()

	fmt.Println(<-output)
	fmt.Println("Горутина 3 — всё")

	/*
		4 вариант

		Вызов os.Exit()
		Exit вызывает выход текущей программы с заданным кодом состояния.
		Как правило, нулевой код означает успех, ненулевой - ошибку.
		Программа завершается немедленно; отложенные функции не выполняются.
	*/

	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			output <- "Горутина 4 — работаем"
		}
	}()

	go func() {
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}()

	for {
		fmt.Println(<-output)
	}

}

func chanCheckStop(ch chan int, wg *sync.WaitGroup) {
	for {
		foo, ok := <-ch
		if !ok {
			println("done")
			wg.Done()
			return
		}
		println(foo)
	}
}

func selectStop(output chan string, quit chan bool) {
	outputString := "Горутина 1: работает"
	for {
		select {
		case output <- outputString:
			//Какие-то действия
		case <-quit:
			return
		}
	}
}
