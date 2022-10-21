package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan struct{})
	for i := 0; i < 3; i++ {
		go func(idx int) {
			select {
			case ch <- (idx + 1) * 2:
				fmt.Println("write: ", (idx+1)*2)
			case value := <-done:
				fmt.Println("exit ", value)
			}
		}(i)
	}

	// get the first result
	fmt.Println(<-ch)
	close(done) // при закрытии канала будут возвращаться дефолтные значения

	time.Sleep(2 * time.Second)
}
