package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	count := 1 
	for ;count<10;count++{
		z -= (z*z - x) / (2*z)
		//fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println("[Sqrt]")
	
	x := 3.0
	fmt.Println(x," --- ", Sqrt(x))

	x = 5.0
	fmt.Println(x," --- ", Sqrt(x))

	x = 7.0
	fmt.Println(x," --- ", Sqrt(x))
}
