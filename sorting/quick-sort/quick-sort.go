package main

import "fmt"

func QuickSort(arr []int, left, right int) {
	if left < right {
		pivot := partition(arr, left, right)
		QuickSort(arr, left, pivot-1)
		QuickSort(arr, pivot+1, right)
	}
}
func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}

func main() {
	arr := []int{1,0,2}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array: ", arr)
}
func QiuckSortalg(arr []int, left, right int) {
	if left < right {
		pivot := partitionAlg(arr, left, right)
		QiuckSortalg(arr, left, pivot-1)
		QiuckSortalg(arr, pivot+1, right)
	}
}
func partitionAlg(arr []int, left, right int) int {
	pivot := arr[right]
	i := left - 1
	for j := left; j < right; j++ {
		if pivot > arr[j] {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}
