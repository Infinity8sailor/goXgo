package main

import "fmt"

// first letter as Capital and its Public
const LoginToken string = "Power"

func main() {

	var username string = "Rein"
	var isLoggedIn bool = false
	var smallVal uint8 = 255
	var floatVal float32 = 244.4444

	fmt.Println(username)
	fmt.Printf("Variable Of type : %T \n", username)

	fmt.Println(isLoggedIn)
	fmt.Printf("Variable Of type : %T \n", isLoggedIn)

	fmt.Println(smallVal)
	fmt.Printf("Variable Of type : %T \n", smallVal)

	fmt.Println(floatVal)
	fmt.Printf("Variable Of type : %T \n", floatVal)

	// Defalut values goes in as 0 [ No Garbage Value ]
	var defaultVal int
	fmt.Println(defaultVal)
	fmt.Printf("Variable Of type : %T \n", defaultVal)

	// Defalut Type goes in as what is assigned and can't be changed later
	var defaulType = "Makima"
	fmt.Println(defaulType)
	fmt.Printf("Variable Of type : %T \n", defaulType)

	// No var type [ but cant be declared outside of method or Globally] :=
	numberOfProtagonists := 3000
	fmt.Println(numberOfProtagonists)
}
