package main

import (
	"fmt"

	"github.com/matheushpr9/algorithms/internal/algorithms"
)

func main() {
	fmt.Println("Quicksort Example")
	arr := []int{3, 6, 6, 6, 6, 8, 10, 1, 2, 1}
	sortedArr := algorithms.Quicksort(arr)
	fmt.Println("Original array:", arr)
	fmt.Println("Sorted array:", sortedArr)
}
