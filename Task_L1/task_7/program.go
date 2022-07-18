package main

import (
	"fmt"
	"sync"
)

// 7. Реализовать конкурентную запись данных в map.

func main() {
	inputMap := make(map[string]int)

	var mutex sync.Mutex // определяем мьютекс

	wg := sync.WaitGroup{}
	wg.Add(3)
	go mapWriting("Первая горутина", inputMap, &mutex, &wg)
	go mapWriting("Вторая горутина", inputMap, &mutex, &wg)
	go mapWriting("Третья горутина", inputMap, &mutex, &wg)

	wg.Wait()

	fmt.Println(inputMap)

}

func mapWriting(goroutineName string, inputMap map[string]int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	mutex.Lock() //блокируем мьютекс так чтобы запись в мапу могла сделать только одна горутина
	for i := 1; i <= 5; i++ {
		inputMap[fmt.Sprintf("%s %d", goroutineName, i)] = i
	}
	mutex.Unlock() //после записи освобождаем мьютекс
	wg.Done()      // Уменьшаем счетчик кол-ва горутин в группе
}
