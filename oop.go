package main

import "fmt"

//композиция структур

type Parent struct {
	name string
}

func (p *Parent) printStructName() {
	fmt.Println("parent")
}
func (p *Parent) printLox() {
	fmt.Println("Lox")
}

type Child struct {
	Parent
}

func (c *Child) printStructName() {
	fmt.Println("child")
}

func main() {
	c := Child{}

	c.Parent.printStructName()
	c.printLox()

	c.name = "ddd"
}
