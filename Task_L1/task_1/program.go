package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

// Action Встраивание методов в структуру с помощью внедрения родительской структуры Human
type Action struct {
	Human
}

type Human struct {
	age    int
	height int
}

func (h Human) getHuman() {
	fmt.Printf("%T age: %d, height: %d", h, h.age, h.height)
}

func main() {
	// Инициализируем объект структуры Human
	human := Human{20, 188}

	// Передаем human в структуру Action
	action := Action{human}

	// у action вызываем функцию getHuman()
	action.getHuman()
}
