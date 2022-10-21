package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "cmd.1"
		ch <- "cmd.2" // не будет обработано
		close(ch)
	}()

	for {
		value, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(value)
	}
}
