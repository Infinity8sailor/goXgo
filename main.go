package main

import "fmt"

// one entry point and reserved keyword
// Case sensitive
func main() {
	fmt.Println("Hello from Rein!")
}

// go run file_name.go [ will just compile and run the file ]
// go help <command> [ For help n all ]

//  go.mod is module
//  Lexer is the program check for the grammer of the language

// Types ( Everything Is Typed )
// String Bool Integer (uint8, int8,....) Floating ( float32, 64)
// Array Slices Maps Stucts Pointers

// Build the program
//  go build [ for default]

// memory management
// There are two methods to do that
// new() => allocates the memory but dont initiate, gives memory address, zeroed storage
// make() => allocates the memory plus initiates it, gives memory address and non-zeroed storage
// GC => Garbage collection happens automatically. but there is a threshould to get in picture

// Pointers
//
