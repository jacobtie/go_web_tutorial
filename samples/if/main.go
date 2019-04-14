package main

import "fmt"

func main() {
	x := 5
	if x == 5 {
		fmt.Println("This code will run if x is equal to 5")
	} else {
		fmt.Println("This code will not run if x is equal to 5")
	}

	y := 24
	if y%5 == 0 {
		fmt.Println("This code will run if y is divisible by 5")
	} else if y%6 == 0 {
		fmt.Println("This code will run if y is divisible by 6")
	} else {
		fmt.Println("This code will run if y is neither divisible by 5 nor 6")
	}

	z := 4
	switch z {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}
}
