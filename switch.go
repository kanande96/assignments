package main

import "fmt"

func main() {
	var i int
	switch i {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("unknown no")
	}
}
