package main

import (
	"bufio"
	"fmt"
	"os"
)

// bufio package provide things for the io
//

func main() {
	welcome := "Welcome Rein"
	fmt.Println(welcome)

	// walrus operation :=
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter The rating for Pizza : ")

	// comma ok syntax || err ok syntax ( catch n err similar)
	input, _ := reader.ReadString('\n')
	fmt.Printf("Thanks : %s", input)

}
