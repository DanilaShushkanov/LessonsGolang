package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	min = 1
	max = 5
)

func main() {
	rand.Seed(time.Now().UnixNano())

	a := rand.Intn(max-min) + min
	b := rand.Intn(max-min) + min

	switch {
	case a > 3:
		fmt.Println(1)
	case b > 3:
		fmt.Println(2)
	default:
		fmt.Println("default")
	}

}
