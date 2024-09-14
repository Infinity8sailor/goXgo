package main

import "fmt"

func main() {

	fmt.Println("Rein is Arraying Here...")

	// You have to define the length of the Array. No Escape
	var myArray [4]string
	myArray[0] = "Mr.Duck"
	myArray[1] = "Mr.Eagle"
	myArray[3] = "Mr.Rook"

	//Unlike other languages it prints the array not needed to iterate
	// Also look at the Enmpty Space for not Initilized Slot
	fmt.Println("myArray Content : ", myArray)
	fmt.Println("Length of Content : ", len(myArray))

	// One more way to initialize the Array
	var pasafistas = [4]string{"Eren", "Annie", "Ben"}

	fmt.Println(pasafistas)
}
