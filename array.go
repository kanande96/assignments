package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 56
	x[1] = 80
	x[2] = 78
	x[3] = 67
	x[4] = 90
	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(total / 5)
}
