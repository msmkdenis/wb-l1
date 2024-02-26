package main

import "fmt"

/*
	Разработать конвейер чисел.
	Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
	после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	go firstProcess(input, firstChan)
	go secondProcess(firstChan, secondChan)

	for v := range secondChan {
		fmt.Println("result from second channel: ", v)
	}
}

func firstProcess(input []int, first chan int) {
	for _, v := range input {
		first <- v
	}
	close(first)
}

func secondProcess(first chan int, second chan int) {
	for v := range first {
		second <- v * 2
	}
	close(second)
}
