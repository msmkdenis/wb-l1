package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
	По истечению N секунд программа должна завершаться.
*/

func main() {
	workTime := time.Second * 3
	ctx, cancel := context.WithTimeout(context.Background(), workTime) // создаем контекст с таймаутом
	defer cancel()

	// канал для обмена сообщениями
	c := make(chan string)

	rnd := rand.NewSource(time.Now().UnixNano())
	go send(ctx, c, rnd) // завершим работу горутины через отмену контексту по таймауту

	wg := &sync.WaitGroup{} // для ожидания завершения горутины, читающей канал
	wg.Add(1)
	go read(c, wg)
	wg.Wait()
}

func send(ctx context.Context, c chan<- string, rnd rand.Source) {
	ticker := time.NewTicker(time.Millisecond * 100) // создаем тикер с интервалом 100 мс для наглядности
	for {
		select {
		case <-ctx.Done():
			fmt.Println("---Контекст завершен по таймауту---")
			close(c)
			return
		case <-ticker.C:
			c <- generateString(rnd)
		}
	}
}

func read(c <-chan string, wg *sync.WaitGroup) {
	for v := range c {
		fmt.Println(v)
	}
	wg.Done()
}

func generateString(rnd rand.Source) string {
	result := make([]byte, 0, 6)
	for i := 0; i < 6; i++ {
		// генерируем случайное число
		randomNumber := rnd.Int63()
		// английские буквы лежат в диапазоне от 97 до 122, поэтому:
		// 1) берем остаток от деления случайного числа на 26, получая диапазон [0,25]
		// 2) прибавляем к полученному числу 97 и получаем итоговый интервал: [97, 122].
		result = append(result, byte(randomNumber%26+97))
	}
	return string(result)
}
