package main

import "fmt"

/*
	Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {
	a := []int{13, 15, -16, 20, 33, 55, -40}
	b := []int{0, 1, 33, 68, 13, 14, 15, 88}
	fmt.Println(intersection(a, b))
}

func intersection(a, b []int) []int {
	var result []int
	m := make(map[int]bool)
	for _, v := range a {
		m[v] = true
	}
	for _, v := range b {
		if _, ok := m[v]; ok {
			result = append(result, v)
		}
	}
	return result
}
