package main

import "fmt"

func main() {
	fmt.Println("Rein is Slicing Here...")
	fmt.Println("Slices is nothing but the Arrays but with abstrations n stuff")
	//  This is Dynamic in Nature
	var Planets = []string{"moon", "phobos", "Eureka"}

	fmt.Printf("Type of Planets : %T \n", Planets)

	Planets = append(Planets, "Charon", "Creta")
	fmt.Println("Planets : ", Planets)

	Planets = append(Planets[1:])
	fmt.Println("Planets : ", Planets)

	// list [inclusive, Exclusive ]
	FewPlanets := append(Planets[1:3])
	fmt.Println("Few Planets : ", FewPlanets)

	// Creating Slice out of Array Using Make
	RealPlanets := make([]int, 4)
	fmt.Println("Real Planets : ", RealPlanets)

	// Using New which just Assigned memory but not initialized
	NewPlanets := new([]int)
	fmt.Println("New Planets : ", NewPlanets)

	// Remove items from the Array
	fmt.Println("Remove the value based on Index.")

	var courses = []string{"React", "Next", "Django", "BootStrap"}
	fmt.Println(courses)
	fmt.Println("Append will be used to do so plus the Spread operator is used.")
	var index = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
