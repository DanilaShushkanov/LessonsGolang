package main

import (
	"fmt"
	"time"
)

func main() {
	forEndlessRange()
}

func nilChannel() {
	var nilChannel chan int
	fmt.Printf("%d, %d \n", len(nilChannel), cap(nilChannel))

	// всегда deadlock, так как некуда писать
	//nilChannel <- 1

	// тоде deadLock Так как нечего читать
	//<-nilChannel

	//закрыть нельзя NIL канал будет паника
	close(nilChannel)
}

func unbufferedChannel() {
	unbuffChan := make(chan int)
	fmt.Printf("%d, %d \n", len(unbuffChan), cap(unbuffChan))

	go func(channel <-chan int) {
		time.Sleep(time.Second / 2)
		value := <-channel
		fmt.Println(value)
	}(unbuffChan)

	unbuffChan <- 228

	go func(channel chan<- int) {
		time.Sleep(time.Second)
		channel <- 123
	}(unbuffChan)

	value := <-unbuffChan
	fmt.Println(value)

	//закрыть можно один раз иначе паника
	close(unbuffChan)

	// читать из закрытого канала можно, всегда будет дефолтное значение для типа
	a := <-unbuffChan
	fmt.Println(a)

	//писать в закрытый канал нельзя будет паника
	//unbuffChan <- 123
}

func bufferedChannel() {
	buffChan := make(chan int, 2)
	fmt.Printf("%d, %d \n", len(buffChan), cap(buffChan))

	//1) два значения могу зщаписать, так как буфер канала 2
	buffChan <- 1
	buffChan <- 2

	go func() {
		time.Sleep(time.Second)
		//4) происходит запись значения в канал
		buffChan <- 3
	}()

	//2) сначала два значения читаются, который я записал в канал
	fmt.Println(<-buffChan)
	fmt.Println(<-buffChan)
	//3) здесь происходит блокировка, так как значения в канале кончились, начинает отрабатывать другая горутина
	//5) можно прочитать, так как в канале появилось значение
	fmt.Println(<-buffChan)

	close(buffChan)

	//при чтении из закрытого канала будут дефолтные значения
	fmt.Println(<-buffChan)

	//при записи в закрытй канал будет паника
	//buffChan <- 123
}

//запусти и все бкудет плюс минус понятно, тут происходжит передача работы с помощью планировщика между рутинами
func forEndlessRange() {
	buffChan := make(chan int, 2)
	numbers := []int{1, 2, 3, 4}

	go func() {
		for _, value := range numbers {
			fmt.Printf("ХОЧУ SEND: %v \n", value)
			buffChan <- value
			fmt.Printf("SEND: %v \n", value)
			fmt.Printf("%d, %d \n", len(buffChan), cap(buffChan))
		}
		close(buffChan)
	}()

	for {
		fmt.Println("я был тут")
		value, ok := <-buffChan
		fmt.Println(value, ok)
		if !ok {
			break
		}
	}
}

func forRange() {
	buffChan := make(chan int, 2)
	numbers := []int{1, 2, 3, 4}

	go func() {
		for _, value := range numbers {
			//2) Приходим сюда, пишем 1 в главную рутину сразу, потом 2 и 3 записываем в буфер
			//3) Хотим записать 4, так как в буфере нет места, блокируемся
			fmt.Printf("ХОЧУ SEND: %v \n", value)
			//5) пишем 4 в канал
			buffChan <- value
			//7) доделываем операции сразу
			fmt.Printf("SEND: %v \n", value)
			fmt.Printf("%d, %d \n", len(buffChan), cap(buffChan))
		}
		close(buffChan)
	}()

	//1) приходим сюда, блокируемся так как нет значений в канале
	fmt.Println("Я был тут")
	//4) Читаем 1,2,3 и потом блокируемся, так как заначений в канале больше нет
	//6) читаем 4 из канала
	for value := range buffChan {
		fmt.Println(value)
	}
}

// тут происходит переключение медлу рутинами и обработка по одному значению, так как буфера у канала нет
func forRangeUnbuffChan() {
	buffChan := make(chan int)
	numbers := []int{1, 2, 3, 4}

	go func() {
		for _, value := range numbers {
			fmt.Printf("ХОЧУ SEND: %v \n", value)
			buffChan <- value
			fmt.Printf("SEND: %v \n", value)
			fmt.Printf("%d, %d \n", len(buffChan), cap(buffChan))
		}
		close(buffChan)
	}()

	fmt.Println("Я был тут")
	for value := range buffChan {
		fmt.Println("unbuff:", value)
	}
}
