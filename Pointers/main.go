package main

import "fmt"

func main() {
	fmt.Println("Rein is here to crack Pointers..")

	// ptr is pointer and it points to the intger type data.
	var ptr *int

	fmt.Println("Pointer : ", ptr) // nil

	myLove := "Rein"
	myLoveAdrr := &myLove

	fmt.Println("my Love address : ", &myLove)
	fmt.Println("& Love : ", myLove)
	fmt.Printf("Type of Love address : %T", &myLove)
	fmt.Println("")
	fmt.Println("& Love Pointer : ", *myLoveAdrr)

	fmt.Println("Setting my new Love to Annie ")

	*myLoveAdrr = "Annie"

	fmt.Println("My new Love is : ", myLove)

}
