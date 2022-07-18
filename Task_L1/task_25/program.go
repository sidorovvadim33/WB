package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		fmt.Println("Внтури рутины начало")
		//time.Sleep(5 * time.Second)
		//OwnSleep(5 * time.Second)
		AnotherSleep(5 * time.Second)
		fmt.Println("Внтури рутины конец")
		wg.Done()
	}()

	wg.Wait()
}

func OwnSleep(d time.Duration) {
	if d <= 0 {
		return
	}

	//Обратный таймер аналогичный заданию 5
	achievableTime := time.Now().Add(d)

	for range time.Tick(1 * time.Millisecond) {
		currentTime := time.Now()
		difference := achievableTime.Sub(currentTime)

		if difference <= 0 {
			break
		}
	}
}

func AnotherSleep(d time.Duration) {
	if d <= 0 {
		return
	}

	//Текущая горутина будет ожидать пока NewTimer не передаст текущее время в канал С
	<-time.NewTimer(d).C // переданное значение никак не обрабатываем просто ожидаем
}
