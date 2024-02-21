package main

import "fmt"

/*
	Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {
	a := []string{"cat", "cat", "dog", "cat", "tree"}
	b := make(map[string]struct{})
	for _, v := range a {
		b[v] = struct{}{}
	}
	output := make([]string, 0, len(b))
	for k := range b {
		output = append(output, k)
	}
	fmt.Println(output)
}
