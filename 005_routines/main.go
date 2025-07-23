package main

import (
	"fmt"
	"runtime"
	"time"
)

func foo() {
	println("This is a goroutine")
}

func main() {
	fmt.Printf("This is operating system of this system:%v\n", runtime.GOOS)
	fmt.Printf("This is architecture of this system:%v\n", runtime.GOARCH)
	fmt.Printf("This is version of Go:%v\n", runtime.Version())
	fmt.Printf("This is number of CPU cores:%v\n", runtime.NumCPU())
	fmt.Printf("This is number of goroutines:%v\n", runtime.NumGoroutine())

	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()
	go foo()

	fmt.Printf("This is number of goroutines before executing foo:%v\n", runtime.NumGoroutine())
	time.Sleep(1 * time.Second)

	fmt.Printf("This is number of goroutines after executing foo:%v\n", runtime.NumGoroutine())
	fmt.Println("This is the end of the main function")
}
