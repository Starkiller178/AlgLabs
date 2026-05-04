package main

import (
	"fmt"
	"lab1/algorithms"
	"lab1/funcfortests"
)

func main() {
	var choice int
	arrN := [4]int{100, 1000, 5000, 10000}
	fmt.Println("ElemIsInArr - 1")
	fmt.Println("SecMaxElem - 2")
	fmt.Println("BinSearch - 3")
	fmt.Println("MultiplyTable - 4")
	fmt.Println("QuickSort - 5")
	fmt.Println("Введите номер алгоритма для тестирования:")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		for i := 0; i < len(arrN); i++ {
			arr := funcfortests.GenerateArr(arrN[i], -10000, 10000)
			t := funcfortests.CalculateTime(func() { algorithms.ElemIsInArr(arr, arr[10]) })
			fmt.Printf("Время работы алгоритма ElemIsInArr для %d элементов: %.5f с\n", arrN[i], t)
		}
	case 2:
		for i := 0; i < len(arrN); i++ {
			arr := funcfortests.GenerateArr(arrN[i], -10000, 10000)
			t := funcfortests.CalculateTime(func() { algorithms.SecMaxElem(arr) })
			fmt.Printf("Время работы алгоритма SecMaxElem для %d элементов: %.5f с\n", arrN[i], t)
		}
	case 3:
		for i := 0; i < len(arrN); i++ {
			arr := funcfortests.GenerateArr(arrN[i], -10000, 10000)
			algorithms.QuickSort(arr, 0, len(arr)-1)
			t := funcfortests.CalculateTime(func() { algorithms.BinSearch(arr, arr[10]) })
			fmt.Printf("Время работы алгоритма BinSearch для %d элементов: %.5f с\n", arrN[i], t)
		}
	case 4:
		for i := 0; i < len(arrN); i++ {
			t := funcfortests.CalculateTime(func() { algorithms.MultiplyTable(arrN[i]) })
			fmt.Printf("Время работы алгоритма MultiplyTable для %d элементов: %.5f с\n", arrN[i], t)
		}
	case 5:
		for i := 0; i < len(arrN); i++ {
			arr := funcfortests.GenerateArr(arrN[i], -10000, 10000)
			t := funcfortests.CalculateTime(func() { algorithms.QuickSort(arr, 0, len(arr)-1) })
			mem := funcfortests.CalculateMemory(func() { algorithms.QuickSort(arr, 0, len(arr)-1) })
			fmt.Printf("Время работы алгоритма QuickSort для %d элементов: %.5f с, Затрачено памяти: %.5f\n", arrN[i], t, mem)
		}
	}
}
