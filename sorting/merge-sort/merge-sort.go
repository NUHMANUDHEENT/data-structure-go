package main

import "fmt"

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return Merge(left, right)
}
func Merge(left, right []int) []int {
	result := []int{}
	var i, j int
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	MergeSort(arr)
	fmt.Println("Sorted array is:", MergeSort(arr))
}
