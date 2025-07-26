package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan int)

	go func() {
		for i := 0; i < 11; i++ {
			fmt.Println("1st function sent ", i, " to channel1")
			channel1 <- i
		}
		close(channel1)
	}()

	go func() {
		for val := range channel1 {
			fmt.Println("2nd function received ", val)
		}
	}()

	time.Sleep(2 * time.Second)
}
