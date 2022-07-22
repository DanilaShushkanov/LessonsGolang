package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	selectManyChannels()
}

func baseSelect() {
	buffChan := make(chan int, 2)

	buffChan <- 1

	// select работает следуюбщим образом
	//сначала анализируются все кейсы, если есть НЕ БЛОКИРУЮЩИЕ, то выполяняется рандомный из них
	//Если все кейсы блокирующие, то выполняется ДЕФОЛТ, если вы его, конечно, написали
	// если не написали, то скорее всего произойдет блокировка и ожидание, до того момента, пока одна из операций не станет НЕБЛОКИРУЮЩЕЙ
	select {
	case num := <-buffChan:
		fmt.Printf("read %v \n", num)
	case buffChan <- 2:
		//здесь закрываю канал, так как если его не закрыть, то произойдет deadlock, из-за блокировки при попытке считывания из канала
		close(buffChan)
		for value := range buffChan {
			fmt.Printf("write: %v \n", value)
		}
	}
}

func selectManyChannels() {
	unbuffChan := make(chan int)
	buffChan := make(chan int, 2)

	buffChan <- 1
	buffChan <- 4

	go func() {
		time.Sleep(time.Second)
		unbuffChan <- 3
	}()

	//Тут все легко и просто TIME.AFTER явялется блокирующей операцией, ровно на столько времени, сколько ему задано
	//Как только это время проходит, а другие операции всё ещё являются блокирующими выполяется этот кейс
	//В таком случае важно понимать, что если есть дефолт, то Time.After работать не будет
	select {
	case num := <-unbuffChan:
		fmt.Println("Print unbuffChan: ", num)
	case buffChan <- 2:
		close(buffChan)
		for value := range buffChan {
			fmt.Printf("Print BuffChan: %v \n", value)
		}
	case <-time.After(time.Second / 2):
		fmt.Println("ВРЕМЯ ВЫШЛО")
	default:
		fmt.Println("DEFAULT")
	}
}

//тут важно понимать, что таймер нужно объявлять за циклом, иначе каждый раз будешь его переписывать
func forRangeSelect() {
	unbuffChan := make(chan int)
	timer := time.After(time.Second)

	go func() {
		defer close(unbuffChan)
		//2) пришли сюда, так как таймер не закончился, то летим в дефолт
		for i := 0; i < 1000; i++ {
			select {
			//4) потом заканчивается таймет, выходим из функции и закрывавем канал, если не закрыть, то будет deadlock
			case <-timer:
				return
			default:
				//3) пишем в канал, потом летим и читаем, опять пишем, опять читаем итд
				time.Sleep(time.Second / 10)
				fmt.Println("ХОЧУ ПИСАТЬ В КАНАЛ")
				unbuffChan <- i
				fmt.Println("ЗАПИСАЛ")
			}
		}
	}()

	//1) пришли сюда, заблочились
	fmt.Println("Я БЫЛ ТУТ")
	for value := range unbuffChan {
		fmt.Println("ОП, ЧИТАЮ")
		fmt.Println(value)
	}
}

func gracefulShutdown() {
	channel := make(chan os.Signal)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)

	timer := time.After(time.Second * 5)

	select {
	case <-timer:
		fmt.Println("Вышли по таймеру")
	case sig := <-channel:
		fmt.Println("ЗАВЕРШИЛИ ", sig)
	}
}
