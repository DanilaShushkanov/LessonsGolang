package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	workerPoolWithTimeout()
}

func baseContext() {
	//базовые контексы
	ctxB := context.Background()
	fmt.Println(ctxB)

	ctxT := context.TODO()
	fmt.Println(ctxT)

	//контекст с отменой
	withCancel, cancel := context.WithCancel(ctxB)
	fmt.Println(withCancel.Err())
	cancel()
	fmt.Println("ОШИБКА ПРИ ОТМЕНЕ КОНТЕКСТА: ", withCancel.Err())

	//хранение значения в контексте
	withValue := context.WithValue(ctxB, "lox", 10)
	fmt.Println("Значение VALUE: ", withValue.Value("lox"))
	fmt.Println(withValue.Err())

	//контекст с дедлайном
	withDeadline, cancel := context.WithDeadline(ctxB, time.Now().Add(time.Second*5))
	defer cancel()
	fmt.Println(withDeadline.Err())
	fmt.Println("Дождались дедлайн: ", <-withDeadline.Done())
	fmt.Println("ОШИБКА ПОСЛЕ ТОГО КАК ДОЖДАЛИСЬ ТАЙМАУТА: ", withDeadline.Err())

	//контекст с таймаутом
	withTimeout, cancel := context.WithTimeout(ctxB, time.Second*2)
	defer cancel()
	fmt.Println("Прошел таймаут: ", <-withTimeout.Done())
	fmt.Println("ОШИБКА ПОСЛЕ ТОГО КАК ДОЖДАЛИСЬ ТАЙМАУТА: ", withTimeout.Err())
}

func workerPoolWithCancel() {
	// создаем контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup

	// создаем два канала, в один будем писать не обработанные значения, во второй будем писать обработанные worker'ами значения
	numbersToProcess, processNumbers := make(chan int, 3), make(chan int, 3)

	//создаем столько рутин, сколько у нас есть виртуальных ядер
	for i := 0; i <= runtime.NumCPU(); i++ {
		//каждую добавляем в waitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			// в каждой рутине свой воркер, который обрабатывает значения
			worker(ctx, numbersToProcess, processNumbers)
		}()
	}

	// записываем 1000 значений в канал необработaнных значений и ЗАКРЫВАЕМ ЕГО
	go func() {
		defer close(numbersToProcess)
		for i := 0; i < 1000; i++ {
			//это пример с выходом по принудительному завершению контекста
			if i >= 500 {
				cancel()
			}
			numbersToProcess <- i
		}
	}()

	// здесь необходимо дождаться выполнения работы всех воркеров, и потом закрываем второй канал, в который они пишут
	go func() {
		wg.Wait()
		close(processNumbers)
	}()

	//запускаем цикл и читаем из второго канала обработанные значения, также считаем, сколько значений успели обработать воркеры
	var counter int
	fmt.Println("я был тут")
	for value := range processNumbers {
		counter++
		fmt.Println(value)
	}

	fmt.Println(counter)
}

func worker(ctx context.Context, numbersToProcess <-chan int, processNumbers chan<- int) {
	for {
		select {
		//это будет блокирующая операция, пока не отменим контекст
		case <-ctx.Done():
			fmt.Println("вышли по контексту")
			return
		// основная операция, станет блокирующей только тогда, когда закроем канал
		case value, ok := <-numbersToProcess:
			if !ok {
				fmt.Println("вышли по закрытию канала")
				return
			}
			time.Sleep(time.Second / 100)
			//возводим в квадрат необработанные значения и записываем во второй канал с обработыннми значениями
			processNumbers <- value * value
		}
	}
}

func workerPoolWithTimeout() {
	// создаем контекст с отменой
	ctx, cancel := context.WithTimeout(context.Background(), time.Second/10)
	defer cancel()
	var wg sync.WaitGroup

	// создаем два канала, в один будем писать не обработанные значения, во второй будем писать обработанные worker'ами значения
	numbersToProcess, processNumbers := make(chan int, 3), make(chan int, 3)

	//создаем столько рутин, сколько у нас есть виртуальных ядер
	for i := 0; i <= runtime.NumCPU(); i++ {
		//каждую добавляем в waitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			// в каждой рутине свой воркер, который обрабатывает значения
			worker(ctx, numbersToProcess, processNumbers)
		}()
	}

	// записываем 1000 значений в канал необработнных значений и ЗАКРЫВАЕМ ЕГО
	go func() {
		defer close(numbersToProcess)
		for i := 0; i < 1000; i++ {
			numbersToProcess <- i
		}
	}()

	// здесь необходимо дождаться выполнения работы всех воркеров, и потом закрываем второй канал, в который они пишут
	go func() {
		wg.Wait()
		close(processNumbers)
	}()

	//запускаем цикл и читаем из второго канала обработанные значения, также считаем, сколько значений успели обработать воркеры
	var counter int
	for value := range processNumbers {
		counter++
		fmt.Println(value)
	}

	fmt.Println(counter)
}
