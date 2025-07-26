package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func makeRequest(i int) int {
	fmt.Println("Making request for", i)

	return i * 5
}

func main() {
	var n int
	fmt.Scan(&n)
	sem := make(chan struct{}, 10)
	results := make(chan int)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sem <- struct{}{}

			result := makeRequest(i)

			results <- result

			fmt.Println("result", result)
			<-sem
		}(i)

	}
	go func() {
		wg.Wait()
		close(results)
	}()

	finalList := []int{}
	for result := range results {
		finalList = append(finalList, result)
	}

	fmt.Println("Final List:", finalList)
}
