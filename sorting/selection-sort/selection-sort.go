package main

import "fmt"

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if i != min {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}
func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	selectionSort(arr)
	fmt.Println("Sorted array is:", arr)
}
