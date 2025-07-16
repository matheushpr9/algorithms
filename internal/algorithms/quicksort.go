package algorithms

func Quicksort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left := []int{}
	right := []int{}
	equal := []int{}

	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		} else {
			equal = append(equal, v)
		}
	}

	return append(append(Quicksort(left), equal...), Quicksort(right)...)
}
