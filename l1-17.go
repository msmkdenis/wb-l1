package main

import "fmt"

/*
	Реализовать бинарный поиск встроенными методами языка.
*/

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // для бин-поиска нужен отсортированный массив

	fmt.Println(brokenSearch(a, 5))
}

func brokenSearch(arr []int, k int) int {
	start := 0
	end := len(arr) - 1

	// Пока начальный индекс меньше или равен конечному
	for start <= end {
		mid := start + (end-start)/2 // Избегаем переполнения

		// Если нашли элемент, возвращаем его индекс
		if arr[mid] == k {
			return mid
		}

		if k >= arr[start] && k < arr[mid] {
			end = mid - 1 // Если элемент меньше, чем средний, то ищем в левой половине
		} else {
			start = mid + 1 // Если элемент больше, чем средний, то ищем в правой половине
		}
	}
	// Если элемент не найден, возвращаем -1
	return -1
}
