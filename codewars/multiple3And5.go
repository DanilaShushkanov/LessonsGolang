package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	fmt.Println(FindMissingLetter([]rune{'a', 'b', 'c', 'd', 'f'}))
}

func Multiple3And5(number int) int {
	if number < 0 {
		return 0
	}
	var count int
	for i := 0; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			count += i
		}
	}

	return count
}

func IsValidWalk(walk []rune) bool {
	if len(walk) != 10 {
		return false
	}

	mapa := make(map[rune]int)
	for i := 0; i < len(walk); i++ {
		if _, ok := mapa[walk[i]]; !ok {
			mapa[walk[i]] = 1
		} else {
			mapa[walk[i]] += 1
		}
	}

	ew := mapa['e'] - mapa['w']
	ns := mapa['n'] - mapa['s']
	return ew == 0 && ns == 0
}

func GetCount(str string) (count int) {
	mapa := map[rune]int{
		'a': 0,
		'e': 0,
		'i': 0,
		'o': 0,
		'u': 0,
	}

	var counter int
	for _, letter := range []rune(str) {
		if _, ok := mapa[letter]; ok {
			counter++
		}
	}

	return counter
}

// надо решать не через рутины
func QueueTime(customers []int, n int) int {
	if len(customers) == 0 {
		return 0
	}

	wg := sync.WaitGroup{}
	toProcess := make(chan int)
	mapa := make(map[string]int)

	for i := 0; i < n; i++ {
		wg.Add(1)
		nameQueue := "queue" + strconv.Itoa(i)
		mapa[nameQueue] = 0
		go func(customers <-chan int, nameQueue string) {
			defer wg.Done()
			for {
				select {
				case v, ok := <-customers:
					if !ok {
						return
					}
					time.Sleep(time.Microsecond * 10 * time.Duration(v))
					mapa[nameQueue] += v
				}
			}
		}(toProcess, nameQueue)
	}

	go func() {
		for _, customer := range customers {
			toProcess <- customer
		}
		close(toProcess)
	}()

	wg.Wait()

	response := 0
	for key, value := range mapa {
		fmt.Printf("%v:%v\n", key, value)
		if response < value {
			response = value
		}
	}

	return response
}

//func DecodeMorse(morseCode string) string {
//	words := strings.Split(strings.TrimSpace(morseCode), "   ")
//	var responseString string
//	for _, word := range words {
//		for _, char := range strings.Split(word, " ") {
//			responseString += MORSE_CODE[char]
//		}
//		responseString += " "
//	}
//
//	return strings.TrimSpace(responseString)
//}

func IsTriangle(a, b, c int) bool {
	return a+b > c && b+c > a && a+c > b
}

type Fighter struct {
	Name            string
	Health          int
	DamagePerAttack int
}

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
	var firtsFighter Fighter
	var secondFighter Fighter
	if fighter1.Name == firstAttacker {
		firtsFighter, secondFighter = fighter1, fighter2
	} else {
		firtsFighter, secondFighter = fighter2, fighter1
	}

	var i int
	for {
		if i%2 == 0 {
			secondFighter.Health -= firtsFighter.DamagePerAttack
		} else {
			firtsFighter.Health -= secondFighter.DamagePerAttack
		}
		i++

		if firtsFighter.Health <= 0 {
			return secondFighter.Name
		} else if secondFighter.Health <= 0 {
			return firtsFighter.Name
		}
	}
}

func add(a, b int) int      { return a + b }
func subtract(a, b int) int { return a - b }
func multiply(a, b int) int { return a * b }
func divide(a, b int) int   { return a / b }

func Arithmetic(a int, b int, operator string) int {
	mapa := map[string]func(int, int) int{
		"add":      add,
		"subtract": subtract,
		"multiply": multiply,
		"divide":   divide,
	}

	return mapa[operator](a, b)
}

func SpinWords(str string) string {
	var spinStr []rune
	for _, word := range strings.Split(str, " ") {
		chars := []rune(word)
		if len(chars) > 5 {
			for i := len(chars)/2 - 1; i >= 0; i-- {
				opp := len(chars) - 1 - i
				fmt.Println(opp, i)
				chars[i], chars[opp] = chars[opp], chars[i]
			}
		}
		spinStr = append(spinStr, chars...)
		spinStr = append(spinStr, ' ')
	}

	return string(spinStr[:len(spinStr)-1])
}

func Int32ToIp(n uint32) string {
	faster64 := int64(n)
	b0 := strconv.FormatInt((faster64>>24)&0xff, 10)
	b1 := strconv.FormatInt((faster64>>16)&0xff, 10)
	b2 := strconv.FormatInt((faster64>>8)&0xff, 10)
	b3 := strconv.FormatInt(faster64&0xff, 10)
	return b0 + "." + b1 + "." + b2 + "." + b3
}

func SquaresInRect(lng int, wdth int) []int {
	if lng == wdth {
		return nil
	}

	slice := make([]int, 0)
	for {
		var x int
		if lng > wdth {
			x = wdth
			lng -= wdth
		} else {
			x = lng
			wdth -= lng
		}

		slice = append(slice, x)
		if wdth == 0 || lng == 0 {
			break
		}
	}

	return slice
}

func FindMissingLetter(chars []rune) rune {
	var findChar rune
	for i := 0; i < len(chars)-1; i++ {
		if chars[i+1]-chars[i] != 1 {
			findChar = chars[i] + 1
			break
		}
	}
	return findChar
}
