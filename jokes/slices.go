package main

import (
	"fmt"
)

func main() {
	testSlices8()
}

//
func testSlices1() {
	a := []string{"a", "b", "c"}
	//тут будет слайс из одного элемента, с 1 элементом
	b := a[1:2]
	//так как слайс А и слайс Б ссылаются на один и тот же базовый массив, при выводе А слайса, там будет q вместо b
	b[0] = "q"
	fmt.Printf("%s\n", a) // что отобразится после вызова?
}

func testSlices2() {
	a := []byte{'a', 'b', 'c'}
	// берем из слайса A 2 элемент(b),у нового слайса будет cap = 2, поэтому мы сможем добавить в слайс еще один элемент без аллокации памяти
	//получится слайс {'b', 'd'}
	b := append(a[1:2], 'd')
	//получится слайс {'z', 'd'}
	b[0] = 'z'
	//получится слайс {'a', 'z', 'd'}
	fmt.Printf("%s\n", a) // что отобразится после вызова?
}

func testSlices3() {
	a := []byte{'a', 'b', 'c'}
	//тут при добавлении 2-ух жлементов мы привысим cap, поэтому вернется слайс, который бует ссылаться на другой базовый массив
	b := append(a[1:2], 'd', 'x')
	b[0] = 'z'
	//поэтому а не поменяется
	fmt.Printf("%s\n", a) // что отобразится после вызова?
}

//ошибка короче
//func testSlices4() {
//	a := []byte{'a', 'b', 'c'}
//	b := string(a)
//	b[0] = 'z'
//	fmt.Printf("%s\n, %T", b[0], b[0]) // что отобразится после вызова?
//}

//тоже ошибка
//func testSlices5() {
//	a := []byte{'a', 'b', 'c'}
//	b := append(a[1:2:1], 'd')
//	b[0] = 'z'
//	fmt.Printf("%s\n", a)
//}

func testSlices6() {
	a := []int{1, 2, 3}
	//3 хуета указывает cap, его можно указать в рамках от количества элементов в слайсе(min), до максимального количества cap, который может получиться от базового слайса(max)
	b := append(a[0:1:2], 4)
	b[0] = 5
	//поэтому тут выведется 5 4 3, 5 появялется из-за подмены, а 4 из-за аппенда
	fmt.Printf("%v\n", a)
}

//прочти про 3 элемент при нарезке слайса
func testSlices7() {
	a := []int{1, 2, 3}
	fmt.Println(a[1:2:2], cap(a[1:2:2]), len(a[1:2:2]))
	b := append(a[1:2:2], 4)
	fmt.Println(b, cap(b), len(b))
	b[0] = 5
	fmt.Printf("%v\n", a)
}

//КРИНГЕ
func testSlices8() {
	//тут короче приколдес высшего уровня
	//сначала ищем элемент с нимуеньшем индексом он будет стоять под своим индексом, все элементы до этого индекса будут заполнены базовыми значениями
	//после элемента с нименьшим индексом будут идти за ним стоящие неиндексирвоанные элементы
	//потом ищем другой нименьший индекс ИТД
	var x = []int{4: 44, 55, 66, 1: 77, 88}
	println(len(x), x[2])
	fmt.Println(x)
}
