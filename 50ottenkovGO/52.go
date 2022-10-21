package main

import "fmt"

type data struct {
	name string
}

func (p *data) print() {
	fmt.Println("name:", p.name)
}

type printer interface {
	print()
}

func main() {
	var22()
}

//func var11() {
//	d1 := data{"one"}
//	d1.print() //ok
//
//	var in printer = data{"two"} // ошибка
//	in.print()
//
//	m := map[string]data{"x": data{"three"}}
//	m["x"].print() //ошибка
//}

func var22() {
	d1 := data{"one"}
	d1.print() //ok

	var in printer = &data{"two"} // ошибка
	in.print()

	m := map[string]*data{"x": &data{"three"}}
	m["x"].print() //ошибка
}
