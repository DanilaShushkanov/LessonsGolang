package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//c([]int{1, 2, 3}...)

	//str1 := "прикказз"
	//str2 := "ккапризз"
	//
	//if isAnagram(str1, str2) {
	//	fmt.Println("Да")
	//} else {
	//	fmt.Println("Нет")
	//}

	//ch1 := make(chan int, 5)
	//ch2 := make(chan int, 5)
	//
	//for i := 0; i < 5; i++ {
	//	ch1 <- i
	//	ch2 <- i * 10
	//}
	//close(ch1)
	//close(ch2)
	//
	//for value := range MergeChannels(ch1, ch2) {
	//	fmt.Println(value)
	//}

	input := make(chan int, 100)

	// sample numbers
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// run goroutine
	go func() {
		for num := range numbers {
			input <- num
		}
		// close channel once all numbers are sent to channel
		close(input)
	}()

	ch1 := DivisionChannels(input)
	ch2 := DivisionChannels(input)

	time.Sleep(3 * time.Second)

	for value := range ch1 {
		fmt.Println("1: ", value)
	}

	for value := range ch2 {
		fmt.Println("2: ", value)
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

func MergeChannels(channels ...<-chan int) <-chan int {
	out := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(len(channels))
		for _, channel := range channels {
			go func(n <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for value := range n {
					out <- value
				}
			}(channel, wg)
		}

		wg.Wait()
		close(out)
	}()

	return out
}

func DivisionChannels(channel <-chan int) <-chan int {
	ch := make(chan int, 100)
	go func() {
		for value := range channel {
			ch <- value
		}
		close(ch)
	}()

	return ch
}
