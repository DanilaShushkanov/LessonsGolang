package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	//a := []int{23, 3, 1, 2}
	//b := []int{6, 2, 4, 23}
	//// [2, 23]
	//fmt.Printf("%v\n", intersection(a, b))
	//a = []int{1, 1, 1}
	//b = []int{1, 1, 1, 1}
	//// [1, 1, 1]
	//fmt.Printf("%v\n", intersection(a, b))

	//for randInt := range randNumGenerator(5) {
	//	fmt.Println(randInt)
	//}

	//a := make(chan int)
	//b := make(chan int)
	//c := make(chan int)
	//
	//go func() {
	//	for _, num := range []int{1, 2, 3} {
	//		a <- num
	//	}
	//	close(a)
	//}()
	//
	//go func() {
	//	for _, num := range []int{20, 10, 30} {
	//		b <- num
	//	}
	//	close(b)
	//}()
	//
	//go func() {
	//	for _, num := range []int{300, 200, 100} {
	//		c <- num
	//	}
	//	close(c)
	//}()
	//
	//for num := range mergeChannels(a, b, c) {
	//	fmt.Println(num)
	//}

	//WorkerPool()

	//fmt.Println(RandInt(5))

	mapa1 := map[string]int{
		"1": 1,
		"2": 2,
	}
	mapa2 := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	fmt.Println(comparisonMap(mapa1, mapa2))
}

func intersection(slice1 []int, slice2 []int) []int {
	counter := make(map[int]int)
	var result []int

	for _, elem := range slice1 {
		if _, ok := counter[elem]; !ok {
			counter[elem] = 1
		} else {
			counter[elem] += 1
		}
	}

	for _, elem := range slice2 {
		if count, ok := counter[elem]; ok && count > 0 {
			counter[elem] -= 1
			result = append(result, elem)
		}
	}

	return result
}

func randNumGenerator(n int) <-chan int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			out <- r.Intn(n)
		}
		close(out)
	}()

	return out
}

func mergeChannels(channels ...<-chan int) <-chan int {
	out := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(len(channels))

		for _, channel := range channels {
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for value := range ch {
					out <- value
				}
			}(channel, wg)
		}

		wg.Wait()
		close(out)
	}()

	return out
}

func sendRead() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x <= 10; x++ {
			fmt.Println("Запись")
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			fmt.Println("Чтение")
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}

func WorkerPool() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	multiplier := func(x int) int {
		return x * 10
	}

	for w := 1; w <= 3; w++ {
		//запустил 3 рутины
		go worker(w, multiplier, jobs, results)
	}

	//1) пришел сюда, записал в буфер 5 значений
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	//for i := 1; i <= numJobs; i++ {
	//	fmt.Println(<-results)
	//}
}

func worker(id int, f func(int) int, jobs <-chan int, results chan<- int) {
	//начнет читать, когда появится первое значение в буфере
	for j := range jobs {
		//переложет реузультат функции в канал result
		results <- f(j)
	}
}

func RandInt(n int) []int {
	mapa := make(map[int]struct{})
	for i := 0; i < n; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		randInt := r.Int()
		if _, ok := mapa[randInt]; ok {
			i--
			continue
		}

		mapa[randInt] = struct{}{}
	}

	slice := make([]int, 0, n)
	for value := range mapa {
		slice = append(slice, value)
	}

	return slice
}

func comparisonMap(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for key, valueA := range a {
		valueB, ok := b[key]
		if !ok || valueA != valueB {
			return false
		}
	}
	return true
}
