package main

import "fmt"

type Flyer interface {
	Fly()
}

type Walker interface {
	Walk()
}

type Swimmer interface {
	Swim()
}

type Ducker interface {
	Flyer
	Walker
	Swimmer
}

type Dog struct {
	Name string
}

func (d *Dog) Walk() {
	fmt.Println("Собака ходит")
}

type Duck struct {
	Name string
}

func (d *Duck) Fly() {
	fmt.Println("Утка летает")
}

func (d *Duck) Walk() {
	fmt.Println("Утка ходит")
}

func (d *Duck) Swim() {
	fmt.Println("Утка плавает")
}

func main() {
	typeAssertionAndPolymorphism()
}

// NilInterface ИНТЕРФЕЙС равен NIL только тогда, когда информация о конкретном типе и его значении равана NIL
func NilInterface() {
	var dI Ducker

	fmt.Printf("%T, %#v\n", dI, dI)
	if dI == nil {
		fmt.Println("ducker is Null")
	}

	var d *Duck
	//d = &Duck{Name: "213"}
	dI = d //появились знание о конкретном типе(если раскомментить строчку выше, то еще и о конкретном значении), поэтому не нил теперь

	fmt.Printf("%T, %#v\n", dI, dI)
	if dI == nil {
		fmt.Println("ducker is Null")
	}
	dI.Walk()
}

func EmptyInterface() {
	var emptyI interface{}
	//x := 123

	//emptyI = x

	fmt.Printf("%T, %#v\n", emptyI, emptyI)
	if emptyI == nil {
		fmt.Println("ducker is Null")
	}
}

func typeAssertionAndPolymorphism() {
	var walker Walker
	fmt.Printf("%T, %#v\n", walker, walker)

	dog := &Dog{Name: "Шарик"}
	duck := &Duck{Name: "Маруся"}

	walker = dog //это абсолтно необязательно, можно просто передавать структуру, которая реализует интерфейс
	polymorphism(walker)
	typeAssertion(walker)

	walker = duck
	polymorphism(walker)
	typeAssertion(walker)
}

func polymorphism(i Walker) {
	i.Walk()
}

func typeAssertion(i Walker) {
	if dog, ok := i.(*Dog); ok {
		fmt.Println("Это собака")
		fmt.Println(dog.Name)
	}

	if duck, ok := i.(*Duck); ok {
		fmt.Println("Это утка")
		fmt.Println(duck.Name)
	}

	switch v := i.(type) {
	case *Dog:
		fmt.Println("ЭТО СОБАКА")
		fmt.Println(v.Name)
	case *Duck:
		fmt.Println("ЭТО УТКА")
		fmt.Println(v.Name)
	default:
		fmt.Println("А ТАКОГО ТИПА НЕТ")
	}
}
