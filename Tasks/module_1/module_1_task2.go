package main

import (
	"fmt"
	"math/rand"
	"time"
)

var que chan int = make(chan int, 10)

func main() {

	defer close(que)

	go consumer()
	producer()

}

func producer() {
	//producer
	for i := 0; ; i++ {
		r := rand.Intn(10000)
		que <- r
		fmt.Printf("sent number: %d\n", r)

		time.Sleep(time.Second)
	}
}

func consumer() {

	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		fmt.Printf("received number: %d\n", <-que)
	}
}
