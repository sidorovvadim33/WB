package main

import "fmt"

/*
Разработать программу, которая в рантайме способна определить
тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

func main() {
	channel := make(chan int)
	fmt.Println("Тип переменной:", getType(channel))
}

//Принимает на вход пустой интерфейс, следовательно можно использовать объект любого типа
func getType(v interface{}) string {
	// %T представление типа значения в Go-синтаксисе
	return fmt.Sprintf("%T", v)
}
