package main

import "fmt"

// go tool vet -shadow .\shdow_varible.go
func main() {
	x := 1
	fmt.Println(x) // выводит 1
	{
		fmt.Println(x) // выводит 1
		x := 2
		fmt.Println(x) // выводит 2
	}
	fmt.Println(x) // выводит 1 (плохо, если нужно было 2)
}
