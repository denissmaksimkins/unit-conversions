package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The proram finds the maximum of three numbers.")
	fmt.Println("Enter three numbers")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	max2 := (a + b + math.Abs(a-b)) / 2
	max := (c + max2 + math.Abs(c-max2)) / 2
	fmt.Println("Max of", a, b, c, "is", max)
}
