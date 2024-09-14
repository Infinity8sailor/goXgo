package main

import "fmt"

func main() {
	fmt.Println("Maps in Go lang.")

	languages := make(map[string]string)

	languages["js"] = "JavaScript"
	languages["py"] = "Python"
	languages["go"] = "GoLang"

	fmt.Println("Maps : ", languages)
	fmt.Println("JS shorts for : ", languages["js"])

	fmt.Println("Deleting the Value in Maps")
	delete(languages, "go")
	fmt.Println("Maps : ", languages)

	// Loops in Maps for no reason
	for key, value := range languages {
		fmt.Printf("Key is %v For value %v \n", key, value)
	}

}
