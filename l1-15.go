package main

import (
	"strings"
)

/*
	К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
	Приведите корректный пример реализации.
*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]rune(v)[:100])
}

func main() {
	someFunc()
}

func createHugeString(l int) string {
	sb := &strings.Builder{}
	for i := 0; i < l; i++ {
		sb.WriteString("a")
	}
	return sb.String()
}
