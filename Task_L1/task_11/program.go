package main

import (
	"fmt"
	"reflect"
)

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {
	firstSlice := []string{"слова", "Какие-то", "записанные", "в массив"}
	secondSlice := []string{"Какие-то1", "слова", "записанные1", "в массив"}

	fmt.Println(getIntersection(firstSlice, secondSlice))
}

func getIntersection(firstSlice []string, secondSlice []string) []reflect.Value {

	mapToCheck := make(map[string]int)

	//Добавляем строки из слайса как ключи в мапу
	for _, elem := range firstSlice {
		mapToCheck[elem] = 0
	}

	//Добавляем строки из 2 слайса как ключи в мапу и увеличиваем значение уже существующих
	for _, elem := range secondSlice {
		if val, ok := mapToCheck[elem]; ok {
			mapToCheck[elem] = val + 1
		}
	}

	//Удаляем из мапы ключи у которых значение не изменилось - элементы не повторлись
	for key, value := range mapToCheck {
		if value == 0 {
			delete(mapToCheck, key)
		}
	}

	//Получаем ключи
	sliceToReturn := reflect.ValueOf(mapToCheck).MapKeys()

	return sliceToReturn
}
