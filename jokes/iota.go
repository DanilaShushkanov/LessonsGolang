package main

import "fmt"

func main() {
	testIOTA3()
}

func testIOTA1() {
	const (
		//int
		p1 = iota // какой тип переменной будет?
		q1 = iota
		r1 = iota
	)

	//0,1,2
	fmt.Printf("%v, %v, %v %T", p1, q1, r1) // что отобразится после вызова?
}

// 0,0,0 так как нет перехода строки и iota не может нормально индексироваться
func testIOTA2() {
	const (
		p2, q2, r2 = iota, iota, iota
	)

	fmt.Println(p2, q2, r2) // что отобразится после вызова?
}

func testIOTA3() {
	const (
		//uint8
		p3 byte = iota
		q3      // какой тип переменной будет? какое будет значение?
		r3
	)

	const (
		p4 = iota
		//все равно продолжит индексироваться, тип int
		q4 // какой тип переменной будет? какое будет значение?
		r4
	)

	fmt.Printf("%T, %v, %T, %v, %v", p3, p3, q3, q3, r3)
}
