package main

import "fmt"

func main() {
	greet()
	fmt.Println("Functions Hello ")

	result := adder(3, 7)
	fmt.Println(" Result of Addition", result)

	proResult := proAdder(3, 7, 10, 20)
	fmt.Println(" Pro Result of Addition", proResult)

	proLiteResult, _ := proLiteAdder(3, 7, 10, 20)
	fmt.Println(" Pro Result of Addition", proLiteResult)

	// Function inside in Functions

}

// Anonymous functions exists and can be called immediately
// func (){
// 	fmt.Println("I am Invincible.")

// }()

func adder(right int, left int) int {
	return right + left
}
func proAdder(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

func proLiteAdder(values ...int) (int, string) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, "Just another fun"
}

func greet() {
	fmt.Println("I am Greeting you ...")
}
