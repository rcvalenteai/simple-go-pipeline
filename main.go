package main

import (
	"fmt"
)

func add(a int, b int) {
	c := a + b
	fmt.Println(c)
}

func main() {
	add(5, 3)
}
