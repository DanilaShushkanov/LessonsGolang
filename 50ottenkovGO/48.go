package main

import "fmt"

func main() {
	var i int = 1

	defer fmt.Println("result =>", func() int { return i * 2 }())
	i++
	//выводит: result => 2 (not ok if you expected 4)
}
