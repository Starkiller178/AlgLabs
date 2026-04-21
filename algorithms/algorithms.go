package algorithms

import (
	"errors"
)

func ElemIsInArr(arr []int, elem int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == elem {
			return true
		}
	}
	return false
}

func SecMaxElem(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("empty array")
	}

	max1 := arr[0]
	max2 := arr[0]

	max2Ex := false
	for i := 0; i < len(arr); i++ {
		if arr[i] > max1 {
			max2 = max1
			max1 = arr[i]
			max2Ex = true
		}

		if arr[i] < max1 && (arr[i] > max2 || !max2Ex) {
			max2 = arr[i]
			max2Ex = true
		}
	}

	if !max2Ex {
		return 0, errors.New("second maximum not found")
	}

	return max2, nil
}

func BinSearch(arr []int, elem int) int {
	low := 0
	hight := len(arr) - 1
	for low <= hight {
		mid := (low + hight) / 2

		if elem == arr[mid] {
			return mid
		}

		if elem > arr[mid] {
			low = mid + 1
		} else {
			hight = mid - 1
		}
	}
	return -1
}

func MultiplyTable(n int) [][]int {
	table := make([][]int, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			table[i][j] = (i + 1) * (j + 1)
		}
	}

	return table
}

func QuickSort(arr []int, low int, hight int) {
	if low < hight {
		pivotInd := partition(arr, low, hight)
		QuickSort(arr, low, pivotInd-1)
		QuickSort(arr, pivotInd+1, hight)
	}
}

func partition(arr []int, low int, hight int) int {
	pivot := arr[low]
	left := low + 1
	right := hight
	for true {
		for left <= right && arr[left] <= pivot {
			left++
		}

		for right >= left && arr[right] >= pivot {
			right--
		}

		if right < left {
			break
		} else {
			temp := arr[left]
			arr[left] = arr[right]
			arr[right] = temp
		}

	}

	temp := arr[low]
	arr[low] = arr[right]
	arr[right] = temp
	return right
}
