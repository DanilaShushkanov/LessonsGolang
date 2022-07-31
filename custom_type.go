package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) printName() {
	fmt.Println(p.Name)
}

func (p *Person) printAge() {
	fmt.Println(p.Age)
}

func main() {
	useCase()
}

func sumWithCustomType() {
	type customType int
	var b customType = 13
	c := 1

	fmt.Println(c + int(b))
}

func getOleg() {
	oleg := Person{Name: "Олег", Age: 11}

	pOleg := &oleg
	fmt.Println((*pOleg).Name)
	fmt.Println(pOleg.Name)
}

// КОМПОЗИЦИЯ/ВСТРАИВАНИЕ

//также можно встративать интерфейсы, но тогда доступ к полям структуры, которая соотвествует этому интерфейсу будет недоступен (можно сделать геттеры и сеттеры)
type Builder struct {
	Person
	Name string
	WorkExperience
}

//func (b *Builder) printName() {
//	fmt.Println(b.Name)
//}

type WorkExperience struct {
	Name string
	Age  int
}

//РАСКОМЕЕНТИ, ТУТ МОНОГО ИНТЕРЕСНОГО
func composition() {
	builder := Builder{
		Person: Person{
			Name: "Степа",
			Age:  30,
		},
		WorkExperience: WorkExperience{
			Name: "Таксист",
			Age:  10,
		},
		Name: "Боб",
	}

	//fmt.Println(builder.Age) // синтаксический сахар (не может достать AGE, так как на одном уровне две структутры имеют это поле)
	fmt.Println(builder.Person.Age)

	fmt.Println(builder.Name) // выведется БОБ, так как это свойство лежит ближе к поверхности композиции

	builder.printName() // тут выведется СТЕПА, так как метод принадлежит структуре PERSON (если раскомментить метод для BUILDER то выведется имя из его структуры)
}

// useCase для композиции с интрефейсом

type BuilderI interface {
	Build()
}

type Building struct {
	BuilderI
	WorkExperience
	Name string
}

type WoodBuilder struct {
	Person
}

func (wb *WoodBuilder) Build() {
	fmt.Printf("%v строит из Дерева", wb.Name)
}

type BrickBuilder struct {
	Person
}

func (bb *BrickBuilder) Build() {
	fmt.Printf("%v строит из Кирпечей", bb.Name)
}

func useCase() {
	buildings := []Building{
		{
			BuilderI: &WoodBuilder{
				Person: Person{
					Name: "СТЕПАН",
					Age:  15,
				},
			},
			WorkExperience: WorkExperience{
				Name: "СТРОИТЕЛЬ ДОМОВ",
				Age:  3,
			},
			Name: "ДЕРЕВЯННЫЙ ДОМ",
		},
		{
			BuilderI: &BrickBuilder{
				Person: Person{
					Name: "АРТЕМ",
					Age:  13,
				},
			},
			WorkExperience: WorkExperience{
				Name: "СТРОИТЕЛЬ ДОМОВ",
				Age:  3,
			},
			Name: "КИРПИЧНЫЙ ДОМ",
		},
	}

	for _, building := range buildings {
		building.Build()
	}
}
