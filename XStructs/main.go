package main

import "fmt"

func main() {
	fmt.Println("Stuct for the Golang")
	// There is no inheritance, parent child Or Super
	rein := User{"Rein", "rein@world.univ", true, 4}
	fmt.Println("User : ", rein)
	fmt.Printf("Name : %v", rein.Name)
	fmt.Printf("All Details : %v \n", rein)

	fmt.Printf("User Status : %v \n", rein.GetStatus())

}

// Has to be in Cap start as going to export same for the Name, Email
// This is how you define
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// This is called method instead of function.we are passing the Struct
// this will copy the user instead of the actual User.
func (user User) GetStatus() bool {
	return user.Status
}
