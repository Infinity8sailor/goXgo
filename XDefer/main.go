package main

import "fmt"

func main() {
	fmt.Println("Defer Started")
	defer fmt.Println(" Im Defer One")
	defer fmt.Println(" Im Defer Two")
	defer fmt.Println(" Im Defer Three")
	MyDefer()
	fmt.Println("")
	fmt.Println("Defer Ended")
}

func MyDefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Printf("%v", i)
	}
}
