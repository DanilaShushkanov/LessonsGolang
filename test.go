package main

import "fmt"

func main() {
	c([]int{1, 2, 3}...)
}

func c(b ...int) {
	fmt.Println(b)
}
