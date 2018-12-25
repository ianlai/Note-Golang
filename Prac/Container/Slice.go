package main

import (
	"fmt"
)

func main() {
	fmt.Println("[Slice]")

	var arr1 [10]int
	var sli1 []int

	sli2 := make([]int, 5, 20)

	fmt.Println(arr1)
	fmt.Println(sli1)
	fmt.Println(sli2)

	printSlice(sli1)
	printSlice(sli2)

	addOneArray(arr1) //not changed
	addOne(sli2)      //will be changed

	fmt.Println("---")

	fmt.Println(arr1)
	fmt.Println(sli1)
	fmt.Println(sli2)

	printSlice(sli1)
	printSlice(sli2)
}

func addOne(T []int) {
	T[0]++
}

func addOneArray(T [10]int) {
	T[0]++
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
