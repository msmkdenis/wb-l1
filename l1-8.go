package main

import "fmt"

/*
	Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func main() {
	fmt.Println(setBit(-1, 63, false))
	fmt.Println(setBit(82, 4, false))
	fmt.Printf("%064b\n", 82)
	fmt.Printf("%064b\n", 66)
}

func setBit(n int64, pos uint, b bool) (int64, error) {
	if pos > 63 {
		return 0, fmt.Errorf("invalid bit position")
	}
	if b {
		n |= 1 << pos
	} else {
		n &= ^(1 << pos)
	}
	return n, nil
}
