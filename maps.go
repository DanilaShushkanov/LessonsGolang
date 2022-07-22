package main

import "fmt"

func main() {
	f()
}

func f() {
	mapNotNil := make(map[string]int)
	mapNotNil["sss"] = 1
	mapNotNil["ddd"] = 2
	mapNotNil["eee"] = 3
	fmt.Printf("%#v \n", mapNotNil)

	changeMap(mapNotNil) // изменится

	fmt.Printf("%#v \n", mapNotNil)
}

func changeMap(mapa map[string]int) {
	mapa["sss"] = 228
}

func mapFunc() {
	//в NIL мапу ничего не добавить будет паника
	//var mapNil map[int]int
	//fmt.Printf("%#v", mapNil)
	//
	//mapNil[1] = 10
	//fmt.Printf("%#v", mapNil)

	mapNotNil := make(map[string]int) // лучше создавать со вторм значеме в MAKE, он выделяет память сразу,
	// иначе при последовательном добавлении может аллоцироваться дофига ненужно памяти

	fmt.Printf("%#v \n", mapNotNil)
	mapNotNil["lox"] = 10
	mapNotNil["neLox"] = 11

	x, ok := mapNotNil["lox"]
	fmt.Printf("Value: %v, OK: %v \n", x, ok)
	k, ok := mapNotNil["loox"]
	fmt.Printf("Value: %v, OK: %v \n", k, ok)
}

type IdName struct {
	Id   int
	Name string
}

func findInSlice() {
	slice := []IdName{
		{1, "Степа"},
		{2, "Миша"},
		{3, "Костя"},
		{4, "Дима"},
		{5, "Даня"},
		{6, "Рома"},
	}
	sought := []string{"Степа", "Дима"}

	mapaFinish1 := make(map[int]*IdName, len(sought))
	for key, currentName := range sought {
		mapaFinish1[key] = findWithoutMap(slice, currentName)
	}

	for _, value := range mapaFinish1 {
		fmt.Println(value)
	}

	//-------------------------- тут поиск через создание мапы, быстрее, чем просто бегать по слайсу
	mapaStart := make(map[string]IdName, len(slice))
	for _, item := range slice {
		mapaStart[item.Name] = item
	}

	mapaFinish2 := make(map[int]*IdName, len(sought))
	for key, currentName := range sought {
		mapaFinish2[key] = findWithMap(mapaStart, currentName)
	}
	for _, value := range mapaFinish2 {
		fmt.Println(value)
	}
}

func findWithoutMap(slice []IdName, currentName string) *IdName {
	for _, IdName := range slice {
		if currentName == IdName.Name {
			return &IdName
		}
	}

	return nil
}

func findWithMap(mapa map[string]IdName, currentName string) *IdName {
	if idName, ok := mapa[currentName]; ok {
		return &idName
	}

	return nil
}
