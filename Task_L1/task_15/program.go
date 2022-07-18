package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

/*
К каким негативным последствиям может привести
данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}


panic: runtime error: slice bounds out of range [:100] with length 2
	justString = v[:100]
в случае если v

*/

var justString string

func someFunc(stringLen int) (string, error) {
	v := createHugeString(1 << 10)

	if len(v) >= stringLen {
		justString = v[:stringLen]
	} else {
		return "", errors.New("строка слишком короткая для введенной длины")
	}

	return justString, nil
}

func createHugeString(num int) string {
	var builder strings.Builder

	for i := 0; i < num; i++ {
		builder.WriteString("i")
	}

	return builder.String()
}

func main() {
	s, err := someFunc(100)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
}
