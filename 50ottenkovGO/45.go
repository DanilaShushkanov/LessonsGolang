package main

import (
	"fmt"
	"sync"
)

type MyLock1 struct {
	sync.Mutex
}

func (ml *MyLock1) print() {
	fmt.Println("1313")
}

//ИЛИ

type MyLock2 sync.Locker //sync.Locker - interface

func main() {
	var ml1 MyLock1
	ml1.Lock()
	ml1.print()
	ml1.Unlock()

	var ml2 MyLock2
	ml2.Lock()
	ml2.Unlock()
}
