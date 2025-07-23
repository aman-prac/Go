package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	wg.Add(25)
	f1 := func(id int) {
		defer wg.Done()
		fmt.Printf("Goroutine %d is running\n", id)
	}
	fmt.Println("Starting 25 goroutines")
	fmt.Printf("Number of Goroutines before WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(1)
	fmt.Printf("Number of Goroutines after 1 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(2)
	fmt.Printf("Number of Goroutines after 2 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(3)
	fmt.Printf("Number of Goroutines after 3 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(4)
	fmt.Printf("Number of Goroutines after 4 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(5)
	fmt.Printf("Number of Goroutines after 5 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(6)
	fmt.Printf("Number of Goroutines after 6 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(7)
	fmt.Printf("Number of Goroutines after 7 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(8)
	fmt.Printf("Number of Goroutines after 8 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(9)
	fmt.Printf("Number of Goroutines after 9 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(10)
	fmt.Printf("Number of Goroutines after 10 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(11)
	fmt.Printf("Number of Goroutines after 11 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(12)
	fmt.Printf("Number of Goroutines after 12 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(13)
	fmt.Printf("Number of Goroutines after 13 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(14)
	fmt.Printf("Number of Goroutines after 14 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(15)
	fmt.Printf("Number of Goroutines after 15 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(16)
	fmt.Printf("Number of Goroutines after 16 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(17)
	fmt.Printf("Number of Goroutines after 17 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(18)
	fmt.Printf("Number of Goroutines after 18 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(19)
	fmt.Printf("Number of Goroutines after 19 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(20)
	fmt.Printf("Number of Goroutines after 20 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(21)
	fmt.Printf("Number of Goroutines after 21 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(22)
	fmt.Printf("Number of Goroutines after 22 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(23)
	fmt.Printf("Number of Goroutines after 23 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(24)
	fmt.Printf("Number of Goroutines after 24 Goroutines in WaitGroup: %d\n", runtime.NumGoroutine())
	go f1(25)
	fmt.Printf("Number of Goroutines after all goroutines in WaitGroup: %d\n", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("All goroutines finished")
	fmt.Println("nowonly goroutine is running", runtime.NumGoroutine())
}
