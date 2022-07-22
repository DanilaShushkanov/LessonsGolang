package main

import "fmt"

type Square struct {
	Side int
}

func (s Square) Perimeter() {
	fmt.Printf("%T %#v \n", s, s)
	fmt.Printf("Perimeter: %d \n", s.Side*4)
}

func (s *Square) Scale(multiplier int) {
	fmt.Printf("%T %#v \n", s, s)
	s.Side *= multiplier
	fmt.Printf("%T %#v \n", s, s)
}

func main() {
	square1 := Square{4}
	pSquare1 := &square1

	square2 := Square{2}

	square1.Perimeter()
	square2.Perimeter()

	pSquare1.Scale(2)    //pSquare.Scale(2)
	square1.Scale(1)     //(&square1).Scale(1)
	pSquare1.Perimeter() //(*pSquare1).Scale(1)
}
