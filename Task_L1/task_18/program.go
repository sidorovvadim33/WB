package main

import "sync"

/*
Реализовать структуру-счетчик,
которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

type Counter struct {
	value int
}

func (c *Counter) increment(mutex *sync.Mutex) {
	mutex.Lock()   // Блокируем мьютекс
	c.value++      // Горутина увеличивает счетчик
	mutex.Unlock() // Освобождаем мьютекс
}

func main() {
	var mutex sync.Mutex // Определяем мьютекс

	wg := sync.WaitGroup{}
	wg.Add(3)
	counter := Counter{0}

	for i := 0; i < 3; i++ {
		go func() {
			for i := 0; i < 3; i++ {
				counter.increment(&mutex)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	println(counter.value)
}
