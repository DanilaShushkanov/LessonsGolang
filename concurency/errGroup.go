package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

func main() {
	errGroup()
}

// по факту с повощью данного подхода, можно делать асинхронный загрузчик
func chanAsPromise() {
	start := time.Now()

	firstResponseChan := makeRequest(1)
	secondResponseChan := makeRequest(2)

	//здесь каждый из каналов будет дожидаться, пока будет получен ответ
	fmt.Println(<-firstResponseChan, <-secondResponseChan)

	fmt.Println(time.Now().Sub(start).Seconds())
}

func makeRequest(numberRequest int) chan string {
	channel := make(chan string, 1)

	go func() {
		time.Sleep(time.Second)
		channel <- fmt.Sprintf("Response number: %v", numberRequest)
	}()
	return channel
}

// просто пример, как можно заменить mutex на каналы
func chanAsMutex() {
	channel := make(chan struct{}, 1)
	var wg sync.WaitGroup
	var counter int

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			//тоже самое, что и обычно, но читабелнее, как мне кажется с использованием Mutex
			channel <- struct{}{}

			counter++

			<-channel
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

// пример с 3 горутинами, здесь мы пытаемся добаиться того, что при ошибке в одной из рутин, будет отменен контекст, от которого зависити выполнение остальных рутин
//Также тут есть проблема с перменной err, в реальной ситуациия, когда несколько рутин могут выдать ошибки, им придется в одну и ту же память (перменная ERR) писать ошибку и может случиться DATARACE
// поэтому, если использовать такой вариант, то нужно еще юзать MUTEX
func withoutErrGroup() {
	ctx, cancel := context.WithCancel(context.Background())
	var err error
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(3)

	// здесь рутина спит 1 секунду имитируя работу, потом мы должны упасть либо в дефолт, либо в первый кейс, если до этого где-то произошла ошибка и контекст отменился
	go func() {
		defer wg.Done()

		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("1 старт")
			time.Sleep(time.Second)
		}
	}()

	// здесь  рутина просто сразу падает в дефолт, потому что мы ничего не ждем и вызываем ошибку и потом отменяем контекст
	go func() {
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("2 старт")
			err = fmt.Errorf("ОШИБКА БЛИНБ")
			cancel()
		}
	}()

	// здесь тоже ничего не ждем и и падаем в дефолт, стартуем и имитируем полезную работу
	go func() {
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("3 старт")
			time.Sleep(time.Second / 2)
			fmt.Println("3 готов")
		}
	}()

	// ждем выполнения 3 рутин
	wg.Wait()

	// если ошибка есть, то выводим её
	// по итогу, отработает 2 и 3 рутины, первая не успеет из-за долго ожидания перед селектом
	// 1 рутина упадет в ctx.Done() и не будет выполняться
	if err != nil {
		fmt.Println(err)
	}
}

// тут юзается errGroup, и она не дает из двух разных рутин выдать ошиюбку и сама отменяет контекст, поэтому здесь не может быть DATARАCE
//
func errGroup() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			return nil
		default:
			fmt.Println("1 старт")
			time.Sleep(time.Second)

			return nil
		}
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
			return nil
		default:
			fmt.Println("2 старт")
			// здесь при получении ошибки, метод g.GO отменит контекст сам и поэтому 1 рутина также не выполнится, как и функции ВЫШЕ
			return fmt.Errorf("ОШИБКА БЛИНБ")
		}
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
			return nil
		default:
			fmt.Println("3 старт")
			time.Sleep(time.Second / 2)
			fmt.Println("3 готов")

			return nil
		}
	})

	//Wait ждет выполнения всех рутин и если одна из них выдала ошибку, то она будет выведена
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
