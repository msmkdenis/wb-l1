package main

import "fmt"

/*
	Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
	Объединить данные значения в группы с шагом в 10 градусов.
	Последовательность в подмножноствах не важна.
	Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func main() {
	input := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	result := make(map[int][]float64)
	for _, v := range input {
		key := int(v/10) * 10
		slice, ok := result[key]
		if !ok {
			slice = []float64{v}
			result[key] = slice
		} else {
			slice = append(slice, v)
			result[key] = slice
		}
	}
	for k, v := range result {
		fmt.Println(k, ":", v)
	}
}
