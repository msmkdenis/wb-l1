package main

import (
	"fmt"
	"time"
)

/*
	Реализовать собственную функцию sleep.
*/

func main() {
	fmt.Println("---------First sleep started: 3 sec duration---------")
	x := time.Now()
	sleepOne(time.Second * 3)
	y := time.Now().Sub(x)
	fmt.Println("Elapsed: ", y)
	fmt.Println("---------Second sleep started: 3 sec duration---------")
	x = time.Now()
	sleepTwo(time.Second * 3)
	y = time.Now().Sub(x)
	fmt.Println("Elapsed: ", y)
}

func sleepOne(duration time.Duration) {
	t := time.Now().Add(duration)
	for time.Now().Before(t) {
	}
}

func sleepTwo(duration time.Duration) {
	<-time.After(duration)
}
