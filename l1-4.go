package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
	Реализовать постоянную запись данных в канал (главный поток).
	Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
	Необходима возможность выбора количества воркеров при старте.

	Программа должна завершаться по нажатию Ctrl+C.
	В результате выбрать и обосновать способ завершения работы всех воркеров.
*/

func main() {
	stop := make(chan os.Signal, 1)                                                       // создаем канал для сигналов
	signal.Notify(stop, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT) // при получении сигнала - отправляем в канал

	// запускаем горутину, в котором закроем канал stop при получении сигнала отмены
	// позволит использовать этот канал в других горутинах
	go func() {
		<-stop
		close(stop)
	}()

	strChan := make(chan string) // создаем канал для строк

	workers := 3            // количество worker-ов, печатающих строку
	wg := &sync.WaitGroup{} // для ожидания завершения всех worker-ов
	wg.Add(workers)
	printer := NewStringPrinter(strChan, workers, wg) // создадим объект для печати
	go printer.Print()                                // запустим в горутине печать с заданным кол-вом worker-ов

	generator := NewStringGenerator(strChan, stop) // создадим объект для генерации
	go generator.Generate()                        // запустим в горутине генерацию строк (можем запустить в main))

	wg.Wait() // ждем завершения всех worker-ов
}

type StringPrinter struct {
	input   <-chan string   // канал для получения строк
	workers int             // количество worker-ов
	wg      *sync.WaitGroup // для ожидания завершения всех worker-ов
	rnd     rand.Source
}

func NewStringPrinter(input <-chan string, workers int, wg *sync.WaitGroup) *StringPrinter {
	return &StringPrinter{input: input, workers: workers, wg: wg}
}

func (r *StringPrinter) Print() {
	for i := 0; i < r.workers; i++ { // запускаем заданное кол-во worker-ов
		go func(i int) {
			for s := range r.input { // читаем из канала, чтение прекратится, как только канал закроется
				fmt.Println("worker", i, " printed: ", s)
			}
			fmt.Println("---Worker", i, " stopped---")
			r.wg.Done() // уменьшаем счетчик ожидаемых worker-ов, код будет выполняться только когда канал строк закроется
		}(i)
	}
}

type StringGenerator struct {
	output chan<- string    // канал для отправки сгенерированных строк
	stop   <-chan os.Signal // канал для получения сигнала отмены
	rnd    rand.Source
}

func NewStringGenerator(output chan<- string, stop <-chan os.Signal) *StringGenerator {
	return &StringGenerator{output: output, stop: stop, rnd: rand.NewSource(time.Now().Unix())}
}

func (r *StringGenerator) Generate() {
	ticker := time.NewTicker(time.Millisecond * 100) // создаем тикер с интервалом 100 мс для наглядности
	for {
		select {
		case _, ok := <-r.stop: // при получении сигнала отмены канал stop закроется
			if !ok { // проверяем закрыт ли канал
				fmt.Println("---Получен сигнал отмены---")
				close(r.output) // закрываем канал, т.к. в этой же функции мы в него пишем.
				return          // выходимаем из горутины
			}
		case <-ticker.C: // по тикеру генерируем строку в канал
			r.output <- r.generateString()
		}
	}
}

func (r *StringGenerator) generateString() string {
	result := make([]byte, 0, 6)
	for i := 0; i < 6; i++ {
		// генерируем случайное число
		randomNumber := r.rnd.Int63()
		// английские буквы лежат в диапазоне от 97 до 122, поэтому:
		// 1) берем остаток от деления случайного числа на 26, получая диапазон [0,25]
		// 2) прибавляем к полученному числу 97 и получаем итоговый интервал: [97, 122].
		result = append(result, byte(randomNumber%26+97))
	}
	return string(result)
}
