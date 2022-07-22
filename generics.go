package main

import "fmt"

func main() {
	typeApproximation()
}

func sumSlice() {
	sliceFloat := []float64{1.0, 2.0, 3.0, 4.0}
	sliceInt := []int64{1, 2, 3, 4}
	//wrongSlice := []interface{}{1, true, struct{}{}}

	//GO сам проверит, подходит ли наш слайс для и использования в функции SUM, если нет то ошибка компиляции
	sumFloat := sum(sliceFloat)
	fmt.Printf("%T, %#v \n", sumFloat, sumFloat)

	//сами указывавем, какого типа слайс мы передаем в функцию
	sumInt := sum[int64](sliceInt)
	fmt.Printf("%T, %#v \n", sumInt, sumInt)

	//Тут нам IDE говорит, что мы передает не тот тип, котоырй ожидается в функции, а ожидается interface{int64 | float64}
	//sum := sum(wrongSlice)
}

func sum[V int64 | float64](numbers []V) V {
	var sum V

	for _, value := range numbers {
		sum += value
	}

	return sum
}

func showContains() {
	type Personn struct {
		Name string
		Age  int64
	}

	sliceInt64 := []int64{1, 2, 3, 4, 5}
	sliceFloat64 := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	sliceStruct := []Personn{
		{"Саша", 10},
		{"Даша", 20},
		{"Паша", 30},
		{"Глаша", 40},
	}

	fmt.Println(contains(sliceInt64, 2))
	fmt.Println(contains(sliceFloat64, 3.0))
	fmt.Println(contains(sliceStruct, Personn{"Пашdа", 30}))

}

// функция, которая сигнализирует, есть ли такое значение в передаваемом слайсе
// можно использовать любые СРАВНИВАЕМЫЕ типы (все кроме слайсов и мапов, и структур, которые содержат несравниваемые типы)
// как раз интерфейс COMPARABLE говорит о том, что тип должен быть СРАВНИВАЕМЫЙ
func contains[V comparable](slice []V, searchValue V) bool {
	for _, value := range slice {
		if value == searchValue {
			return true
		}
	}

	return false
}

// принимает значения ЛЮБОГО ТИПА, превращает в слайс и принтит
func show[V any](entities ...V) {
	fmt.Println(entities)
}

//-------- TYPE PARAMETERS --------
//Юзать либо в чистую int64 | float64 (UNION)
// либо comparable
// Либо создать свой интерфейс, в котором НЕ БУДЕТ методов

type Number64 interface {
	int64 | float64
}

func sumUnionInterfaceSlice() {
	sliceFloat := []float64{1.0, 2.0, 3.0, 4.0}
	sliceInt := []int64{1, 2, 3, 4}
	//wrongSlice := []interface{}{1, true, struct{}{}}

	//GO сам проверит, подходит ли наш слайс для и использования в функции SUM, если нет то ошибка компиляции
	sumFloat := sumUnionInterface(sliceFloat)
	fmt.Printf("%T, %#v \n", sumFloat, sumFloat)

	//сами указывавем, какого типа слайс мы передаем в функцию
	sumInt := sumUnionInterface[int64](sliceInt)
	fmt.Printf("%T, %#v \n", sumInt, sumInt)

	//Тут нам IDE говорит, что мы передает не тот тип, котоырй ожидается в функции, а ожидается interface{int64 | float64}
	//sum := sum(wrongSlice)
}

func sumUnionInterface[V Number64](numbers []V) V {
	var sum V

	for _, value := range numbers {
		sum += value
	}

	return sum
}

// Numbers64 -------- ОБОЩЕННЫЙ ТИП
// Numbers64 данный тип - это слайс, и тип его элементов ограничен interface'ом NUMBER64
type Numbers64[T Number64] []T

func sumUnionInterfaces() {
	// когда создается переменная с обощенным типом, то необходимо указывать тип значений, какой будет в ней лежать
	var sliceInt64 Numbers64[int64]
	sliceInt64 = append(sliceInt64, []int64{1, 2, 3, 4, 5}...)

	sumInt := sum[int64](sliceInt64)
	fmt.Printf("%T, %#v \n", sumInt, sumInt)

	//так тоже можно создавать перменные с обощенным типом
	sliceFloat64 := Numbers64[float64]{1.0, 2, 3, 4.0}

	sumFloat := sum(sliceFloat64)
	fmt.Printf("%T, %#v \n", sumFloat, sumFloat)
}

//--- ПРИБЛИЖЕНИЕ ТИПОВ ----

type Number32 interface {
	~int32 | float32
}

type CustomInt32 int32

func (ci CustomInt32) IsPositive() bool {
	return ci > 0
}

func typeApproximation() {
	sliceCustomInt32 := []CustomInt32{1.0, 2.0, 3.0, 4.0, 5.0}
	//это первый вариант, можно перконвертировать каждое значение,(долго да и зачем)
	sliceCastedInt32 := make([]int32, 0, len(sliceCustomInt32))
	for _, value := range sliceCustomInt32 {
		sliceCastedInt32 = append(sliceCastedInt32, int32(value))
	}

	sumCustomInt32 := sumApproximation(sliceCustomInt32)
	sumCastedInt32 := sumApproximation(sliceCastedInt32)

	//  если не делать приблежение типа, то будет ошибка компиляции
	//значение будет тип CustomInt32
	fmt.Printf("%T, %v \n", sumCustomInt32, sumCustomInt32)
	// значение будет типа int32
	fmt.Printf("%T, %v \n", sumCastedInt32, sumCastedInt32)
}

func sumApproximation[V Number32](numbers []V) V {
	var sum V

	for _, value := range numbers {
		sum += value
	}

	return sum
}
