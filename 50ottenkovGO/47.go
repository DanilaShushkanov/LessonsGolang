package main

import (
	"fmt"
	"time"
)

func main() {
	var4()
}

func var1() {
	slice := []string{"1", "2", "3"}

	for _, v := range slice {
		go func() {
			fmt.Println(v)
		}()
	}

	time.Sleep(3 * time.Second)
}

func var2() {
	slice := []string{"1", "2", "3"}

	for _, v := range slice {
		go func(n string) {
			fmt.Println(n)
		}(v)
	}

	time.Sleep(3 * time.Second)
}

//------------------------------------------------------------------------------

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

func var3() {
	data := []field{{"one"}, {"two"}, {"three"}}

	for _, v := range data {
		go func(f field) {
			f.print()
		}(v)
	}

	time.Sleep(3 * time.Second)
}

func var4() {
	data := []*field{{"one"}, {"two"}, {"three"}}

	for _, v := range data {
		go v.print()
	}

	time.Sleep(3 * time.Second)
}
