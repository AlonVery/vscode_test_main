package main

import "fmt"

// BubbleSort сортирует срез чисел в порядке возрастания
func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Меняем местами элементы
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	if len(arr) < 4 {
		fmt.Print("hello world")
	}
	return arr
}
func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original array:", arr)
	sortedArr := BubbleSort(arr)
	fmt.Println("Sorted array:", sortedArr)

	fmt.Println("testing changed with git command in terminal")
}
