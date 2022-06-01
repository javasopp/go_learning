package main

import (
	"fmt"
	"time"
)

func workder(id int, c chan int) {
	for {
		fmt.Printf("Worker %d received %d\n", id, <-c)
	}
}

func chanDemo() {
	c := make(chan int)
	go workder(1, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	time.Sleep(time.Millisecond)
	//var c chan int // c == nil
}

func main() {
	chanDemo()
}
