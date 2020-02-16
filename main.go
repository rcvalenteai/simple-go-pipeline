package main

import (
	"fmt"
)

func add(a int, b int) int {
	c := a + b
	fmt.Println(c)
	return c
}

func main() {
	add(5, 3)
}
