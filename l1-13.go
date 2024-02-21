package main

import "fmt"

/*
	Поменять местами два числа без создания временной переменной.
*/

func main() {
	//Если имеется в виду поменять значения в переменных, то:
	x, y := 5, 10
	x = x ^ y
	y = y ^ x
	x = x ^ y
	fmt.Println(x, y)

	x = x + y
	y = x - y
	x = x - y
	fmt.Println(x, y)

	//Если имеется в виду поменять значения в слайсе:
	a := []int{1, 5}
	a[0], a[1] = a[1], a[0]
	fmt.Println(a)
}
