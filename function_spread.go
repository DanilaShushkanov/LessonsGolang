package main

import "fmt"

type MyError struct{}

func (MyError) Error() string {
	return "MyError!"
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	// Каков будет вывод программы и почему?
	var err *MyError
	errorHandler(err)

	err = &MyError{}
	errorHandler(err)
}

//func main() {
//	b := []int{1, 2, 3, 4, 5}
//	a(b...)
//}
//
//func a(nums ...int) { // в виде []int
//	for _, value := range nums {
//		fmt.Println(value)
//	}
//}
