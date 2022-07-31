package main

import "fmt"

func main() {
	//c([]int{1, 2, 3}...)

	str1 := "прикказз"
	str2 := "ккапризз"

	if isAnagram(str1, str2) {
		fmt.Println("Да")
	} else {
		fmt.Println("Нет")
	}
}

func c(b ...int) {
	fmt.Println(b)
}

func isAnagram(str1 string, str2 string) bool {
	fmt.Println(str1)
	fmt.Println(str2)

	rune1 := []rune(str1)
	rune2 := []rune(str2)

	fmt.Println(rune1)
	fmt.Println(rune2)
	if len(rune1) == len(rune2) {
		mapa := make(map[rune]int, len(rune1))
		for _, elem := range rune1 {
			if _, ok := mapa[elem]; !ok {
				mapa[elem] = 1
			} else {
				mapa[elem] += 1
			}
		}

		for _, elem := range rune2 {
			if _, ok := mapa[elem]; !ok {
				return false
			} else {
				mapa[elem] -= 1
				if el, _ := mapa[elem]; el <= 0 {
					delete(mapa, elem)
				}
			}
		}

		if len(mapa) == 0 {
			return true
		}
	}

	return false
}

func Merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {

}
