package main

import (
	"fmt"
	"sync"
)

func main() {
	// This is an example of race condition in GO
	incrementor := 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	ws := 1000
	wg.Add(ws)
	for i := 0; i < ws; i++ {
		go func() {
			mu.Lock()
			v := incrementor
			v++
			incrementor = v

			wg.Done()
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Final value of incrementor:", incrementor)
}
