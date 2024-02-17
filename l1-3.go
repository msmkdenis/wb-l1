package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

func main() {
	a := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	fmt.Println("---Второй способ - отправляем результат в канал и читаем из него в другой горутине---")

	res := make(chan int)
	go func() { // в отдельной горутине запустим чтение из канала и выведем результат
		var sum int
		for r := range res {
			sum += r
		}
		println(sum)
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
