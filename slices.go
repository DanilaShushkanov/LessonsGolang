package main

import (
	"fmt"
)

func main() {
	reverseSlice()
}

func filterSlice() {
	slice := []int{1, 2, 21, 1, 2, 32, 13, 13, 123, 13, 123, 13, 1, 1}

	n := 0
	for _, x := range slice {
		if x > 5 {
			slice[n] = x
			n++
		}
	}
	slice = slice[:n]
	fmt.Printf("%v, %v, %v \n", slice, len(slice), cap(slice))
}

func reverseSlice() {
	slice := []int{1, 2, 3, 4, 5}

	for i := len(slice)/2 - 1; i >= 0; i-- {
		opp := len(slice) - 1 - i
		fmt.Println(opp, i)
		slice[i], slice[opp] = slice[opp], slice[i]
	}
	fmt.Printf("%v, %v, %v \n", slice, len(slice), cap(slice))
}

func getSlice() {
	arr := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%T, %#v \n", arr, arr)

	//slice от array
	slice := arr[1:3]
	fmt.Printf("%v, %v, %v \n", slice, len(slice), cap(slice))

	arr[1] = 99 //изменилось значение и в слайсе и в массиве, так как этот массив является базовым для салайса
	fmt.Printf("%T, %#v \n", arr, arr)
	fmt.Printf("%v, %v, %v \n", slice, len(slice), cap(slice))

	//сделал реслайзинг
	newSlice := slice[:1]
	fmt.Printf("%v, %v, %v \n", newSlice, len(newSlice), cap(newSlice))

	// изменилость во всех слайсах и массиве
	newSlice[0] = 2
	fmt.Printf("%T, %#v \n", arr, arr)
	fmt.Printf("%v, %v, %v \n", slice, len(slice), cap(slice))
	fmt.Printf("%v, %v, %v \n", newSlice, len(newSlice), cap(newSlice))

	//такой аппенд изменяет сразу в трех местах, добавляя значение в NewSlice, изменяее значение в Slice и Array
	newSlice = append(newSlice, 22)
	fmt.Printf("%T, %#v \n", arr, arr)
	fmt.Printf("%v, %v, %v \n", slice, len(slice), cap(slice))
	fmt.Printf("%v, %v, %v \n", newSlice, len(newSlice), cap(newSlice))

	// такой аппенд обновляет значения в базовом массиве и докидывает в newSlice, базовывй массив не пересоздается, так как cap хватает, если добавить еще элемент
	//то базовый массив пересоздастся и переменная Array изменяться не будет, а newSlice увеличит свое капасити и получет элементы в себя
	newSlice = append(newSlice, 33, 44)
	fmt.Printf("%T, %#v \n", arr, arr)
	fmt.Printf("%v, %v, %v \n", slice, len(slice), cap(slice))
	fmt.Printf("%v, %v, %v \n", newSlice, len(newSlice), cap(newSlice))
}

func copySlice() {
	destination := make([]int, 0, 2)
	source := []int{1, 2, 3}

	fmt.Println("Copied: ", copy(destination, source))
	fmt.Printf("%#v, %v, %v \n", destination, len(destination), cap(destination))

	//копируется столько значений сколько длина слайса в который копируешь
	destination = make([]int, 2, 3)
	fmt.Println("Copied: ", copy(destination, source))
	fmt.Printf("%#v, %v, %v \n", destination, len(destination), cap(destination))

	//или максимальное количество элементов, котоыре есть в слайсе из которого копируем
	destination1 := make([]int, 3, 3)
	source1 := []int{1}
	fmt.Println("Copied: ", copy(destination1, source1))
	fmt.Printf("%#v, %v, %v \n", destination1, len(destination1), cap(destination1))

	//чтобы скпировать всё, используй len()
	source2 := []int{1, 2, 3}
	destination2 := make([]int, len(source2))
	fmt.Println("Copied: ", copy(destination2, source2))
	fmt.Printf("%#v, %v, %v \n", destination2, len(destination2), cap(destination2))

	//В НИЛОВЫЙ СЛАЙС НИЧЕГО НЕ ВСТАВИТСЯ, ТАК КАК ДЛИНА 0
}

func deleteElement() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("%#v, %v, %v \n", slice, len(slice), cap(slice))

	i := 2
	//самый простой способ
	slice = append(slice[:i], slice[i+1:]...)
	fmt.Printf("%#v, %v, %v \n", slice, len(slice), cap(slice))
}

func passToFunction() {
	slice := []int{1, 2}
	fmt.Printf("%v, %v, %v \n", slice, len(slice), cap(slice))

	changeSlice(slice) // все изменится так как слайс содержит указатель на базовывмй массив
	fmt.Printf("%v, %v, %v \n", slice, len(slice), cap(slice))

	slice = append(slice, 1)
	fmt.Printf("%v, %v, %v \n", slice, len(slice), cap(slice))

	slice = appendSlice(slice)
	fmt.Printf("%v, %v, %v \n", slice, len(slice), cap(slice))
}

func changeSlice(slice []int) {
	slice[0] = 10
}

func appendSlice(slice []int) []int {
	slice = append(slice, 3)
	fmt.Printf("%v, %v, %v \n", slice, len(slice), cap(slice))

	return slice
}

func appendToSlice() {
	x := make([]int, 0, 2)
	fmt.Printf("%v, %v, %v \n", x, len(x), cap(x))
	x = append(x, 1, 2, 3, 4, 5, 6, 7)
	fmt.Printf("%v, %v, %v \n", x, len(x), cap(x))

	x = append(x, 6)
	fmt.Printf("%v, %v, %v \n", x, len(x), cap(x))

	//----------------------- разложения слайса на отдельные переменные
	firstSlice := []int{1, 2, 3, 4, 5}
	secondSlice := []int{6, 7, 8, 9}

	firstSlice = append(firstSlice, secondSlice...)

	fmt.Printf("%v, %v, %v \n", firstSlice, len(firstSlice), cap(firstSlice))
}

//  нужно чтобы совпадала длина массива и слайса
func convertToArrayPointer() {
	initialSlice := []int{1, 2}
	fmt.Printf("%#v, %v, %v \n", initialSlice, len(initialSlice), cap(initialSlice))

	array := (*[2]int)(initialSlice)
	fmt.Printf("%#v, %v, %v \n", array, len(array), cap(array))
}

func arrayFunc() {
	arr := [4]int{1, 2, 3, 4}

	fmt.Printf("%T, %#v \n", arr, arr)
	changeArray(&arr)
	fmt.Printf("%T, %#v \n", arr, arr)
}

func changeArray(arr *[4]int) {
	arr[3] = 11
}
