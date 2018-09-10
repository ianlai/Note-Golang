package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i := 0
	c1 := 0
	c2 := 1
	return func() int {
		i = c1
		c1 = c1 + c2
		c2 = i
		return i
	}
}

func main() {
	fmt.Println("[Fibonacci]")
	
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
