package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a,b, значение которых > 2^20.
*/

func main() {
	var x int64
	fmt.Scan(&x)
	//Для работы с большим числами использовалась библиотека "math/big"
	a := big.NewInt(x)

	fmt.Scan(&x)
	b := big.NewInt(x)

	c := big.NewInt(0)
	f := big.NewFloat(0.0)

	//Устанавливает знаечние с равному a*b
	fmt.Printf("Умножение = %d \n", c.Mul(a, b))

	aFloat := big.NewFloat(0.0).SetInt(a)
	bFloat := big.NewFloat(0.0).SetInt(b)

	//Устанавливает z в округленное значение x/y
	fmt.Printf("Деление = %f \n", f.Quo(aFloat, bFloat))

	//Устанавливает значение а равному а+b
	fmt.Printf("Сумма = %d \n", c.Add(a, b))

	//Устанавливает значение а равному а-b
	fmt.Printf("Вычитание = %d \n", c.Sub(a, b))
}
