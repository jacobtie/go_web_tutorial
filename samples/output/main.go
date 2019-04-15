package main

import "fmt"

func main() {
	// Prints Hello without a line break
	fmt.Print("Hello")
	// Prints World! with a line break
	fmt.Println("World!")
	// Prints the number 5 with a line break using %d
	fmt.Printf("%d is a number\n", 5)
	// This also works with a variable
	num := 6
	fmt.Printf("%d is also a number\n", num)
	// Use %s to print strings
	line := "Go is a neat language"
	fmt.Printf("%s that we can use to build servers\n", line)
}
