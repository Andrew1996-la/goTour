package simpeslices

import "fmt"

func Simpeslices() {

	/*
		это массив. Он всегда имеет заданную длинну
	*/
	arr := [4]int{1, 3, 4, 5}
	
	/*
		это слайс он по объявлению похож на массив но не имеет установленной длинны
		можно модифицировать
	*/
	// mySkills := []string{
	// 	"golang",
	// 	"linux",
	// 	"docker",
	// }

	/*
		Так же мы можем создвать слайсы из массивов
	*/
	sliceFromArr := arr[0:2]

	fmt.Println(sliceFromArr)


	/*
		проверка на пустоту слайса осуществляется через len
	*/
	var emtySlice []int
	fmt.Println(len(emtySlice) == 0)

	/*
		так же мы можем изначально алоцировать участок в памяти под слайс
		с помощью make
	*/

	makeSlice := make([]int, 10, 20)
	newList := make([]int, len(makeSlice), cap(makeSlice))

	/*
		Для копирования одного слайсв в дроугой используется copy
	*/
	copy(newList, makeSlice)

	fmt.Println("makeSlice", len(makeSlice), cap(makeSlice))
	fmt.Println("newList", len(newList), cap(newList))

	// testSlice := make([]int, 5, 5);

	// aa :=  Append(testSlice, 99)
	// fmt.Println("aa", testSlice)
}

func Append(list []int, elem int) []int {
	newList := make([]int, len(list) + 1, cap(list)+1)
	copy(newList, list)

	newList[len(list)] = elem
	return newList
}