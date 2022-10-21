package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	testGoroutines6()
}

// deadlock так как ничего записать и считать из nilChannel нельзя
func testGoroutines1() {
	var ch chan int
	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- (idx + 1) * 2
		}(i)
	}
	fmt.Println("result:", <-ch)
	time.Sleep(2 * time.Second)
}

// что мы увидим в stdout?
// увидим cmd1 cmd2
func testGoroutines2() {
	ch := make(chan string)
	go func() {
		//2)читаю из канала 1 запись
		//4) читаю из канала 2 запись
		for m := range ch {
			fmt.Println("processed:", m)
		}
	}()
	//1) прихожу сюда, записываю в канал
	ch <- "cmd.1"
	//3) пишу в канал 2 запись
	ch <- "cmd.2"
}

// здесь всегда будет выводиться последнее из цикла, так как в рутины не передаются значения,
// и V в рутине вычисляется при выводе через fmt.Println(), а не при создании рутины
func testGoroutines3() {
	data := []string{"one", "two", "three"}
	for _, v := range data {
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(3 * time.Second)
}

// какая последняя рутина отработает то значение и будет в num
func testGoroutines4() {
	var num int

	for i := 0; i < 10000; i++ {
		go func(i int) {
			num = i
		}(i)
	}
	fmt.Printf("NUM is %d", num)
}

// ошибка, конкурентно
// без синхронизации в map писать нельзя
func testGoroutines5() {
	dataMap := make(map[string]int)
	mu := sync.Mutex{}
	//ch := make(chan struct{}, 1)

	for i := 0; i < 10000; i++ {
		go func(d map[string]int, num int) {

			//ch <- struct{}{}
			mu.Lock()
			d[fmt.Sprintf("%d", num)] = num
			mu.Unlock()
			//<-ch

		}(dataMap, i)
	}
	time.Sleep(5 * time.Second)
	fmt.Println(len(dataMap))
}

// до какой итерации успеем дойти в цикле
// то и будет записано
func testGoroutines6() {
	runtime.GOMAXPROCS(1)

	x := 0
	go func(p *int) {
		for i := 1; i <= 20000000000; i++ {
			*p = i
		}
	}(&x)

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("x = %d.\n", x)
}
