package main

import (
	"fmt"
	"strings"
)

/*
	Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
	Функция проверки должна быть регистронезависимой.

	Например:
	abcd — true
	abCdefAaf — false
	aabcd — false
*/

func main() {
	fmt.Println(isUnique("abcd"))
	fmt.Println(isUnique("abCdefAaf"))
	fmt.Println(isUnique("aabcd"))
	fmt.Println(isUnique("bB"))
}

func isUnique(s string) bool {
	s = strings.ToLower(s)
	set := make(map[rune]struct{})
	for _, r := range s {
		if _, ok := set[r]; ok {
			return false
		}
		set[r] = struct{}{}
	}
	return true
}
