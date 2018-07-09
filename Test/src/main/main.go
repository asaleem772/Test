package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	for i, v := range []int{1, 2, 3, 4} {
		fmt.Println(i, v)

	}
}
