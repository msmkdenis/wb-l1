package main

import (
	"fmt"
	"sync"
)

/*
	Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	a := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	fmt.Println("---Первый способ - на каждый элемент массива запускаем горутину и сразу выводим на печать---")
	wg.Add(len(a))
	for i := 0; i < len(a); i++ {
		go func(i int) {
			defer wg.Done()
			println(a[i] * a[i])
		}(i)
	}
	wg.Wait()

	fmt.Println("---Второй способ - отправляем результат в канал и читаем из него в другой горутине---")
	res := make(chan int)
	go func() { // в отдельной горутине запустим чтение из канала
		for r := range res {
			println(r)
		}
	}()

	wg.Add(len(a))
	for i := 0; i < len(a); i++ {
		go func(i int) {
			defer wg.Done()
			res <- a[i] * a[i]
		}(i)
	}
	wg.Wait()  // ждем окончания работы цикла
	close(res) // закрываем канал т.к. мы уверены, что больше писать в него не будем
}
