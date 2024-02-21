package main

import (
	"fmt"
	"sync"
)

/*
	Реализовать конкурентную запись данных в map.
*/

func main() {
	fmt.Println("---Реализация с использованием стандартной map и sync.RWMutex---")
	simpleMap := make(map[int]int)
	simpleMapWrite(simpleMap)
	fmt.Println(simpleMap)
	fmt.Println("---Реализация с использованием sync.Map---")
	var syncMap sync.Map
	syncMapWrite(&syncMap)
	syncMap.Range(func(k, v interface{}) bool {
		fmt.Println("key:", k, ", val:", v)
		return true // if false, Range stops
	})
}

func simpleMapWrite(m map[int]int) {
	wg := sync.WaitGroup{}
	mu := sync.RWMutex{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			m[i] = i
		}(i)
	}
	wg.Wait()
}

func syncMapWrite(m *sync.Map) {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(i, i)
		}(i)
	}
	wg.Wait()
}
