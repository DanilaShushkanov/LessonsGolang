package main

import "fmt"

func main() {
	getElMap()
}

func f() {
	mapNotNil := make(map[string]int)
	mapNotNil["sss"] = 1
	mapNotNil["ddd"] = 2
	mapNotNil["eee"] = 3
	fmt.Printf("%#v \n", mapNotNil)

	changeMap(mapNotNil) // изменится

	fmt.Printf("%#v \n", mapNotNil)
}

func changeMap(mapa map[string]int) {
	mapa["sss"] = 228
}

func mapFunc() {
	//в NIL мапу ничего не добавить будет паника
	//var mapNil map[int]int
	//fmt.Printf("%#v", mapNil)
	//
	//mapNil[1] = 10
	//fmt.Printf("%#v", mapNil)

	mapNotNil := make(map[string]int) // лучше создавать со вторм значеме в MAKE, он выделяет память сразу,
	// иначе при последовательном добавлении может аллоцироваться дофига ненужно памяти

	fmt.Printf("%#v \n", mapNotNil)
	mapNotNil["lox"] = 10
	mapNotNil["neLox"] = 11

	x, ok := mapNotNil["lox"]
	fmt.Printf("Value: %v, OK: %v \n", x, ok)
	k, ok := mapNotNil["loox"]
	fmt.Printf("Value: %v, OK: %v \n", k, ok)
}

type IdName struct {
	Id   int
	Name string
}

func findInSlice() {
	slice := []IdName{
		{1, "Степа"},
		{2, "Миша"},
		{3, "Костя"},
		{4, "Дима"},
		{5, "Даня"},
		{6, "Рома"},
	}
	sought := []string{"Степа", "Дима"}

	mapaFinish1 := make(map[int]*IdName, len(sought))
	for key, currentName := range sought {
		mapaFinish1[key] = findWithoutMap(slice, currentName)
	}

	for _, value := range mapaFinish1 {
		fmt.Println(value)
	}

	//-------------------------- тут поиск через создание мапы, быстрее, чем просто бегать по слайсу
	mapaStart := make(map[string]IdName, len(slice))
	for _, item := range slice {
		mapaStart[item.Name] = item
	}

	mapaFinish2 := make(map[int]*IdName, len(sought))
	for key, currentName := range sought {
		mapaFinish2[key] = findWithMap(mapaStart, currentName)
	}
	for _, value := range mapaFinish2 {
		fmt.Println(value)
	}
}

func findWithoutMap(slice []IdName, currentName string) *IdName {
	for _, IdName := range slice {
		if currentName == IdName.Name {
			return &IdName
		}
	}

	return nil
}

func findWithMap(mapa map[string]IdName, currentName string) *IdName {
	if idName, ok := mapa[currentName]; ok {
		return &idName
	}

	return nil
}

func getElMap() {
	//есть 4 способа создать мапу
	var mapa1 map[int]int //это nil мапа, в неё ничего записать не полчится, считать тоже, будет паника (СПОСОБ ГОВНО)
	fmt.Println(mapa1)
	//mapa1[1] = 2
	var mapa2 = new(map[int]int) //это указатель на nil мапу, что по факту несет за собо те же проблемы, что и способ выше (СПОСОБ ГОВНО)
	fmt.Println(mapa2)
	var mapa3 = map[int]int{} //это уже жизнеспособный вариант инциализации мапы оне не будети нил
	fmt.Println(mapa3)
	//тут все супер так как не nil мапа
	mapa3[1] = 2
	var mapa4 = make(map[int]int, 4) //хороший вариант, можно указать capacity, вторым значением в make()
	fmt.Println(mapa4)
	mapa4[1] = 2

	//считывание из мапы
	//нельзя взять указатель на адресс элемента, так как это небоезопасно из-за эвакуации данных, которы производится, когда среднее значение элементов
	// во всех bucket становится больше 6.5
	//el := &mapa4[1]

	// получение значения происходит по следующему алгоритму
	// 1) Вычисляется значение хеша от ключа
	// 2) По хешу и размеру bucket понимаем в каком bucket лежит искомый элемент
	// 3) Вычисляется дополнительный хеш из старших 8 битов старого хеша
	// 4) В найденом bucket поочередно сравниваем дополнительный хеш каждого элемента с дополнительным хешом ключа
	// 5) Если Хеши не совпали, то переходим в следующий bucket, ссылка не него хранится в этом bucket (односвязный список)
	// 6) Если мы дошли до bucket в котором нет ссылки на слеющий bucket и все еще не нашли элемент, то возвращается дефолт того типа, которое значение в мапе
	// 7) Если дополнительные хеши совпали, то определяется где лежит в памяти ключ, подозреваемы как искомый, сравнимаем равен ли он тому, что запросили ЕСЛИ ДА, то возврщаем значение
	el := mapa4[1]
	fmt.Println(el)
}
