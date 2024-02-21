package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
	По завершению программа должна выводить итоговое число
*/

type CounterAtomic struct {
	counter atomic.Int64
	c       int
}

func main() {
	count := &CounterAtomic{}
	countAtomic(count, 1000)
	fmt.Println("atomic counter (should be 1000): ", count.counter.Load())
	fmt.Println("race counter (should be 1000): ", count.c)
}

func countAtomic(count *CounterAtomic, times int) {
	var wg sync.WaitGroup
	for i := 0; i < times; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.counter.Add(1)
			count.c++
		}()
	}
	wg.Wait()
}
