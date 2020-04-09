package main

import (
	"fmt"
)

const (
	first = iota + 6
	second
)

func main() {
	fmt.Println(first, second)
}
