package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке
уникальные (true — если уникальные, false etc). Функция проверки должна быть
регистронезависимой.

Например:
	abcd — true
	abCdefAaf — false
	aabcd — false
*/

func main() {
	var s string
	for {
		fmt.Println("\nВведите строку")
		fmt.Scan(&s)

		fmt.Println("Проверка мапой:", checkSymbolsWithMap(s))
		fmt.Println("Проверка через сплит и сортировку:", checkSymbolsWithStrings(s))
	}
}

func checkSymbolsWithMap(s string) bool {
	symbolsMap := make(map[int32]int32, len(s))
	/*
		Перебираем строку и добавляем символы в map,
		в случае если значение в map уже добавлено
		сразу возвращаем false
	*/
	for _, elem := range s {
		if _, ok := symbolsMap[elem]; ok {
			return false
		} else {
			symbolsMap[elem] = 1
		}
	}
	return true
}

func checkSymbolsWithStrings(s string) bool {
	symbols := strings.Split(s, "") // Создаем слайс букв из строки
	sort.Strings(symbols)           // сортируем: таким образом повторяющиеся буквы будут находится рядом

	for i := 1; i < len(symbols)-1; i++ { // Перебор слайса начиная со второго элемента
		if symbols[i] == symbols[i-1] { // Если предыдущий символ равен текущему сразу возврщаем false
			return false
		}
	}
	return true
}
