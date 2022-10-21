package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	testData1()
	fmt.Println("---------------------------------------")
	testData2()
}

type MyData struct {
	One int    `json:"one"`
	two string `json:"two"`
}

// тут прикол с обработкой структур функциями из других пакетов
// то есть поля, которые НЕЭКСПОРТИРУЕМЫЕ будут либо заполняться стандартными значениями для своего типа
// либо не будут отображены, как например в случае с json.Marshal()
func testData1() {
	in := MyData{1, "two"}
	fmt.Printf("%#v\n", in) // что отобразится после вызова?
	encoded, _ := json.Marshal(in)
	fmt.Println(string(encoded)) // что отобразится после вызова?
	var out MyData
	json.Unmarshal(encoded, &out)
	fmt.Printf("%#v\n", out) // что отобразится после вызова?
}

// здесь прикол в том, что в локальной области (в данном случае цикл for1) всегда записывается указатель на V и в каждом ключе слайса будет записан указательна одну и ту же ячейку памяти(в ней хранится V)
// а в итоге, последним значем в V будет 4, поэтому выведутся все 4
func testData2() {
	a := []int{1, 2, 3, 4}
	result := make([]*int, len(a))
	for i, v := range a {
		result[i] = &v
	}
	for _, u := range result {
		fmt.Printf("%d ", *u) // что отобразится после вызова?
	}
}
