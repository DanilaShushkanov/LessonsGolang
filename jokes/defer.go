package main

import "fmt"

type toProcess struct {
	toReturn int
}

func (t *toProcess) process() {
	t.toReturn = 1
	fmt.Println("2", t.toReturn)
}

func (t *toProcess) testFunc() int {
	//перед отображением в main выполнится эта функция и отобразится 1
	defer t.process()
	t.toReturn = 2
	//вернется 2 и отобразится в main
	return t.toReturn
}

func main() {
	var t = toProcess{}
	fmt.Println("1", t.testFunc()) // что отобразится после вызова?
}
