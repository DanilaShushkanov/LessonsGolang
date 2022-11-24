package main

import (
	"fmt"
	"sync"
)

func main() {
	wgWithChannel()
}

func mutexWithChannel() {
	var counter int
	var wg sync.WaitGroup
	ch := make(chan struct{}, 1)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- struct{}{}
			counter++
			<-ch
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

func wgWithChannel() {
	var counter int
	n := 1000
	wg := make(chan struct{}, n)
	ch := make(chan struct{}, 1)

	for i := 0; i < n; i++ {
		go func(k int) {
			ch <- struct{}{}
			counter++
			<-ch

			wg <- struct{}{}
		}(i)
	}

	for i := 0; i < n; i++ {
		<-wg
	}

	fmt.Println(counter)
}
