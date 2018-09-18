package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

/*
WordCount function
*/
func WordCount(s string) map[string]int {
	sArr := strings.Fields(s)
	m := make(map[string]int)

	for _, value := range sArr {
		m[value] = m[value] + 1
	}
	return m
}

func main() {
	fmt.Println("[Word Count]")

	s1 := "this is a new pen from a man with another pen"
	s2 := "aaa bb a ccc c bb bb aa a ccc bb aaa bb"
	map1 := WordCount(s1)
	map2 := WordCount(s2)

	fmt.Println("s1:")
	for k, v := range map1 {
		fmt.Println(k, v)
	}
	fmt.Println()
	fmt.Println("s2:")
	for k, v := range map2 {
		fmt.Println(k, v)
	}
	wc.Test(WordCount)
}
