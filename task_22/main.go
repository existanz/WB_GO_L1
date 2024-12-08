/*
Разработать программу, которая перемножает, делит, складывает, вычитает
две числовых переменных a и b, значение которых > 2^20.
*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	// на самом деле 2^20 = 1048576 не такое большое число
	// максимальное число int64 = 9223372036854775807 = 2^63 - 1
	// что вполне достаточно для операций с числами > 2^20
	// для ещё больших чисел можно использовать float64
	// максимальное число float64 = 1.7976931348623157e+308 = 2^1023 - 1
	// кроме того обычно для больших чисел используют пакет big.Int

	a := big.NewInt(1 << 57)
	b := big.NewInt(1 << 42)
	c := new(big.Int)

	fmt.Println("a:", a, "b:", b)

	c.Mul(a, b)
	fmt.Println("Multiplication:", c)

	c.Div(a, b)
	fmt.Println("Division:", c)

	c.Add(a, b)
	fmt.Println("Addition:", c)

	c.Sub(a, b)
	fmt.Println("Subtraction:", c)

}
