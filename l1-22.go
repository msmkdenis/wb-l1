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
	var a, b big.Int
	fmt.Scan(&a, &b)
	c := sum(&a, &b)
	fmt.Println(c)

}

func sum(a, b *big.Int) *big.Int {
	return a.Add(a, b)
}

func sub(a, b *big.Int) *big.Int {
	return a.Sub(a, b)
}

func mul(a, b *big.Int) *big.Int {
	return a.Mul(a, b)
}

func div(a, b *big.Int) (*big.Int, bool) {
	if b.Cmp(big.NewInt(0)) == 0 {
		return nil, false
	}
	return a.Div(a, b), true
}
