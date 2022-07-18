package main

import "fmt"

/*
Имеется последовательность строк -
(cat, cat, dog, cat, tree)
создать для нее собственное множество.
*/

func main() {
	var slice = []string{"cat", "name", "cat", "dog", "cat", "tree", "name"}

	fmt.Println(createSet(slice))
}

func createSet(strings []string) []string {
	stringsMap := make(map[string]bool)
	// Перебираем слайс и добавляем в мапу, дубликатов не будет тк ключи уникальны
	for _, elem := range strings {
		stringsMap[elem] = true
	}

	uniqueStrings := make([]string, 0, len(stringsMap))
	for key := range stringsMap {
		// все ключи из map в слайс
		uniqueStrings = append(uniqueStrings, key)
	}

	return uniqueStrings
}
