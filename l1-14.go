package main

import (
	"fmt"
	"time"
)

/*
	Разработать программу, которая в рантайме способна определить тип переменной:
	int, string, bool, channel из переменной типа interface{}.
*/

func main() {
	in := make(chan interface{})
	go getType(in)
	in <- 5
	in <- "hello"
	in <- true
	in <- make(chan interface{})
	close(in)
	time.Sleep(1 * time.Second)
}

func getType(in chan interface{}) {
	for c := range in {
		switch c.(type) {
		case int:
			fmt.Println("int: ", c)
		case string:
			fmt.Println("string: ", c)
		case bool:
			fmt.Println("bool: ", c)
		case chan interface{}:
			fmt.Println("channel: ", c)
		default:
			fmt.Println("unknown type")
		}
	}
}
