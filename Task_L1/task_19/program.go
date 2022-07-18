package main

/*
Разработать программу, которая переворачивает
подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.
*/

func main() {
	s := "главрыба"

	println(reverseStringWithRange(s))

	println(reverseStringWithRune(s))
}

func reverseStringWithRune(s string) string {
	asRune := []rune(s)
	//Перебираем массив рун с двух сторон к центру, меняя местами крайние буквы
	for i, j := 0, len(asRune)-1; i < j; i, j = i+1, j-1 {
		asRune[i], asRune[j] = asRune[j], asRune[i]
	}
	return string(asRune)
}

func reverseStringWithRange(s string) (stringToReturn string) {
	//Перебор строки: ставим каждую новую бувку в начало итоговой
	for _, s := range s {
		stringToReturn = string(s) + stringToReturn
	}
	return stringToReturn
}
