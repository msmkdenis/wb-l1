package main

import (
	"context"
	"fmt"
	"time"
)

/*
	Реализовать все возможные способы остановки выполнения горутины.
	Остановка - выход из горутины (заершение работы), не блокирование.
*/

func main() {
	// отмена контекста через ручной вызов cancel
	ctxFirst, firstCancel := context.WithCancel(context.Background())
	defer firstCancel()
	go first(ctxFirst)
	go stopFirst(firstCancel)

	time.Sleep(time.Second * 1)

	// отмена контекста по deadline-у
	ctxSecond, secondCancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*1))
	defer secondCancel()
	go second(ctxSecond)

	time.Sleep(time.Second * 1)

	// отмена контекста по timeout-у
	ctxThird, thirdCancel := context.WithTimeout(context.Background(), time.Second*1)
	defer thirdCancel()
	go third(ctxThird)

	time.Sleep(time.Second * 1)

	// закрытие сигнального stop-канала
	stop := make(chan struct{})
	go fourth(stop)
	go stopFourth(stop)

	time.Sleep(time.Second * 1)

	// истечение времени
	after := time.After(time.Second * 1)
	go fifth(after)

	time.Sleep(time.Second * 1)

	// истечение таймера
	timer := time.NewTimer(time.Second * 1)
	go six(timer)

	time.Sleep(time.Second * 1)

	// по тикеру
	ticker := time.NewTicker(time.Second * 1)
	go seventh(ticker)

	time.Sleep(time.Second * 2)
	fmt.Println("exit main")
}

func first(ctx context.Context) {
	ticker := time.NewTicker(time.Millisecond * 200) // создаем тикер с интервалом 100 мс для наглядности
	for {
		select {
		case <-ctx.Done():
			fmt.Println("first stopped by calling cancel, exit")
			return
		case <-ticker.C:
			fmt.Println("first is running")
		}
	}
}

func stopFirst(cancel context.CancelFunc) {
	time.Sleep(time.Second * 1)
	cancel()
}

func second(ctx context.Context) {
	ticker := time.NewTicker(time.Millisecond * 200) // создаем тикер с интервалом 100 мс для наглядности
	for {
		select {
		case <-ctx.Done():
			fmt.Println("second stopped by deadline, exit")
			return
		case <-ticker.C:
			fmt.Println("second is running")
		}
	}
}

func third(ctx context.Context) {
	ticker := time.NewTicker(time.Millisecond * 200) // создаем тикер с интервалом 100 мс для наглядности
	for {
		select {
		case <-ctx.Done():
			fmt.Println("third stopped by timeout, exit")
			return
		case <-ticker.C:
			fmt.Println("third is running")
		}
	}
}

func fourth(stop chan struct{}) {
	ticker := time.NewTicker(time.Millisecond * 200) // создаем тикер с интервалом 100 мс для наглядности
	for {
		select {
		case _, ok := <-stop:
			if !ok {
				fmt.Println("fourth stopped by closing stop channel")
				return
			}
		case <-ticker.C:
			fmt.Println("fourth is running")
		}
	}
}

func stopFourth(stop chan struct{}) {
	time.Sleep(time.Second * 1)
	close(stop)
}

func fifth(stop <-chan time.Time) {
	ticker := time.NewTicker(time.Millisecond * 200) // создаем тикер с интервалом 100 мс для наглядности
	for {
		select {
		case <-stop:
			fmt.Println("fifth stopped by timer.After, exit")
			return
		case <-ticker.C:
			fmt.Println("fifth is running")
		}
	}
}

func six(stop *time.Timer) {
	ticker := time.NewTicker(time.Millisecond * 200) // создаем тикер с интервалом 100 мс для наглядности
	for {
		select {
		case <-stop.C:
			fmt.Println("six stopped by timer, exit")
			return
		case <-ticker.C:
			fmt.Println("six is running")
		}
	}
}

func seventh(stop *time.Ticker) {
	ticker := time.NewTicker(time.Millisecond * 200) // создаем тикер с интервалом 100 мс для наглядности
	for {
		select {
		case <-stop.C:
			fmt.Println("seventh stopped by ticker, exit")
			return
		case <-ticker.C:
			fmt.Println("seventh is running")
		}
	}
}
