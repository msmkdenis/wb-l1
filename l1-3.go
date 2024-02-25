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

	res := make(chan int)
	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(len(a))
		for i := 0; i < len(a); i++ {
			go func(i int) {
				defer wg.Done()
				res <- a[i] * a[i]
			}(i)
		}
		wg.Wait()
		close(res) // закрываем канал т.к. мы уверены, что больше писать в него не будем
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() { // в отдельной горутине запустим чтение из канала и выведем результат
		var sum int
		defer wg.Done()
		for r := range res { // пока не будет закрыт канал - чтение будет продолжаться
			sum += r
		}
		fmt.Println(sum)
	}()
	wg.Wait() // дождем вывода результата
}
