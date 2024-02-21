package main

import (
	"fmt"
	"slices"
)

/*
	Удалить i-ый элемент из слайса
*/

func main() {
	slice := []int{1, 2, 3, 4, 5}
	slice = slices.Delete(slice, 2, 3)
	fmt.Println(slice)

	slice = []int{1, 2, 3, 4, 5}
	slice = anotherDelete(slice, 2)
	fmt.Println(slice)

	anotherSlice := []string{"a", "b", "c", "d", "e"}
	anotherSlice = anotherDelete(anotherSlice, 2)
	fmt.Println(anotherSlice)

	slice = []int{1, 2, 3, 4, 5}
	temp := slice[:2]
	slice = append(temp, slice[3:]...)
	fmt.Println(slice)
}

func anotherDelete[S interface{ ~[]E }, E any](s S, i int) S {
	out := s[:i]
	out = append(out, s[i+1:]...)
	return out
}
