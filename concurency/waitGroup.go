package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	writeWithoutMutex()

}

//проблема в том, что main завершается быстрее, чем выполняются другие 10 горутин
func withoutWait() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i + 1)
	}

	fmt.Println("exit")
}

//тут все хорошо, но если используется общая память в горутинах, то нужен мьютекс
func withWait() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i + 1)
		}(i)
	}

	wg.Wait()
	fmt.Println("exit")
}

// медленно, но выводятся всё что нужно
func writeWithoutConcurrent() {
	start := time.Now()
	var counter int

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Second / 100)
		counter++
	}

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

// сильно быстрее, но есть датарейсы в counter
func writeWithoutMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second / 100)
			counter++
		}()
	}

	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

// чуть медленне чем предыдущий вариант, но теперь выводятся все записи
func writeWithMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second / 100)
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

//------- MUTEX VS RWMUTEX
//в данной функции юзается простой мьютекс, поэтому при блокировке горутиной, которая читает данные, получается так, что другие горутины
// которые пишут и читают блокируются
// данная функция щзанимает больше времени, чем следующая
func writeReadWithMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(100)
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Second / 100)

			mu.Lock()
			fmt.Println(counter)
			mu.Unlock()
		}()

		go func() {
			defer wg.Done()
			time.Sleep(time.Second / 100)

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

//здесь используется RWMUTEX, и горутины которые читают, блокируют только те горутины которые пишут
// а горутины которые пишут, блокируют все остальные горутины
// данная фнукция отраюатывает быстрее, но здесь при чтении  могуцт прочитаться одни и те же номера  counter'a
func writeReadWithRWMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mu sync.RWMutex

	wg.Add(100)
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Second / 100)

			mu.RLock()
			_ = counter
			mu.RUnlock()
		}()

		go func() {
			defer wg.Done()
			time.Sleep(time.Second / 100)

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}
