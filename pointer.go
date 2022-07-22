package main

import "fmt"

func main() {
	a := 2
	square(a)
	fmt.Println(a)
	pointerSquare(&a)
	fmt.Println(a)
}

func pointerSquare(a *int) {
	*a *= *a
}

func square(a int) {
	a *= a
}
