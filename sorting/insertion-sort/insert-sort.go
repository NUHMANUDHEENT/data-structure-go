package main

import "fmt"

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		value := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > value {
			arr[j+1] = arr[j]
			j -= 1
		}
		arr[j+1] = value
	}
}
func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	InsertionSort(arr)
	fmt.Println("Sorted array is:", arr)
}
