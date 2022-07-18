package main

import (
	"fmt"
	"regexp"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func main() {
	s := "snow dog sun"
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	// С помощью регулярных выражений разделяем строку по пробелам в массив
	splittedString := regexp.MustCompile(`\s+`).Split(s, -1)

	builder := strings.Builder{}
	for i := len(splittedString) - 1; i >= 0; i-- {
		//Конкатенация строк
		builder.WriteString(splittedString[i] + " ")
	}

	//Удаляем лишний пробел в конце строки
	return strings.TrimSpace(builder.String())
}
