package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

type Worker struct {
	number int
}

func (worker *Worker) print(chanelInt chan int) {
	for {
		fmt.Println("Worker", worker.number, "значение из канала:", <-chanelInt)
		time.Sleep(600 * time.Millisecond)
	}
}

func main() {
	chanelInt := make(chan int)
	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	fmt.Print("Введите кол-во воркеров: ")
	var workerCount int
	_, err := fmt.Scanln(&workerCount)
	// Если при вводе поймали ошибку или количество Воркеров меньше 1 завершаем программу
	if err != nil || workerCount < 1 {
		return
	}

	// Создаем объекты структуры Worker с номерами и запускаем у каждого в отдельной горутине функцию print
	for i := 1; i <= workerCount; i++ {
		worker := Worker{i}
		go worker.print(chanelInt)
	}

	result := 1

	for {
		//Оператор select ожидает нескольких операций отправки или получения одновременно
		select {
		case chanelInt <- result:
			result++
		case <-c:
			fmt.Println("Program stopped")
			os.Exit(0)
		}
	}
}
