package main

import (
	"fmt"
)

func main() {
	deferValues()
}

func panicRecover() {
	defer func() {
		//recover выдаст значение panic
		fmt.Println(recover())
	}()

	// когда случается паника программа старается быстрее завершится, раскручивается stackTrace до главной функции
	panic("exit")

	// recover БЕЗ defer не сработает никогда
	fmt.Println(recover())

	// это уже не выведется, так как паника
	fmt.Println(123)
}

func deferValues() {
	// тут всё окей просто в цикле добавляются defer, значение i, запоминается сразу в них, после завершения функциия они выводятся в обратном порядке
	for i := 0; i < 10; i++ {
		if i > 5 {
			break
		}
		defer fmt.Println(i)
	}

	//здесь всё не окей, потому что здесь значение i сразу не вычисляется, а defer func добавляется в defer стек с переменной I
	//и I будет вычислятся после RETURN
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}

	//-------- как исправить такое поведение --------
	// делаем праметр и прокидываем его внутрь функции
	for i := 0; i < 10; i++ {
		defer func(k int) {
			fmt.Println(k)
		}(i)
	}

	// создаем локальную переменную  и она уже в defer будет запоминаться верно
	for i := 0; i < 10; i++ {
		k := i
		defer func() {
			fmt.Println(k)
		}()
	}
}

func deferredFunc() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println(sum(2, 3))

	fmt.Println("finish")
}

//сначала складываются значения в перменную SUM, потом происходит RETURN и уже потом в defer func() значение SUM домножается на 2 и выводится
func sum(x, y int) (sum int) {
	defer func() {
		sum *= 2
	}()

	sum = x + y
	return
}

func printNumbers() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}

// перенеси в main, чтобы всё работало как надо
func f() {
	fmt.Println("start")

	go printNumbers()

	// задать количество виртуальных ядер, на которых будет выполняться работа рутин
	// 1 ядро = 1 горутина
	//runtime.GOMAXPROCS(1)
	//количество виртуальных ядер
	//fmt.Println(runtime.NumCPU())

	// переключение на другую горутину с помощью планировщика
	//time.Sleep(time.Second)
	//принудительное переключение на другую горутину
	//runtime.Gosched()

	fmt.Println("finish")
}
