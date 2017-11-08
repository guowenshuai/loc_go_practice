package main

import (
	"fmt"
	"time"
)

func counting(c chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		c <- i
	}
	close(c)
}

func main() {
	msg := "Staring main"
	fmt.Println(msg)
	bus := make(chan int)
	msg = "staring a gofunc"
	go counting(bus)
	for count := range bus {
		fmt.Println("count:", count)
	}
}
