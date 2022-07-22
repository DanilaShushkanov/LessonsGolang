package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	addValue()
}

// тоже самое, что и обычно, только с пакетом atomic работает чуть быстрее, чем с mutex
// но в нем есть только простые типы
func addAtomic() {
	start := time.Now()

	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func addMutex() {
	start := time.Now()

	var counter int64
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()
			counter++
			mu.Unlock()

		}()
	}

	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func storeLoadSwap() {
	var value int64

	//атомарно достает значение
	fmt.Println(atomic.LoadInt64(&value))
	//атомарно вставляет значение
	atomic.StoreInt64(&value, 1)
	fmt.Println(atomic.LoadInt64(&value))

	//атомарно заменяет значение и возвращает предыдущее
	fmt.Println(atomic.SwapInt64(&value, 2))
	fmt.Println(atomic.LoadInt64(&value))

	//атомарно заменяет значение, если оно было равно OLD
	//если было равно OLD, то возврщает true иначе false
	fmt.Println(atomic.CompareAndSwapInt64(&value, 2, 3))
	fmt.Println(atomic.CompareAndSwapInt64(&value, 2, 3))
}

// useCase, когда необходимо только одной рутиной изменить значение какой-то переменной
func compareAndSwap() {
	var value int64
	var wg sync.WaitGroup

	// запускается 100 рутин
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()

			// здесь одна из рутин сможет поменять значение с 0 на 1 и вернутся true
			// все остальные будут менять с 1 на 1 и уже условие не выполнится и вернется false
			if !atomic.CompareAndSwapInt64(&value, 0, 1) {
				return
			}

			// сюда попадет только одна рутина, которая смогла заменить значение
			fmt.Println("Рутина с этим нмоером заменила ", k)
		}(i)
	}

	wg.Wait()
}

// также существует тип atomic.Value с методами, как и функции в пакете atomic
func addValue() {
	//тип interface{}
	var value atomic.Value

	value.Store(1)
	fmt.Println(value.Load())

	fmt.Println(value.Swap(2))

	fmt.Println(value.CompareAndSwap(2, 3))
	fmt.Println(value.CompareAndSwap(2, 3))
}
