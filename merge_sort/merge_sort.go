package merge_sort

/**
1. Create a slice that size is equal to left node plus right node
*/
func merge(left []int, right []int) []int {
	size, leftIdx, rightIdx := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for idx := 0; idx < size; idx++ {
		if leftIdx >= len(left) && rightIdx < len(right) {
			slice[idx] = right[rightIdx]
			rightIdx++
		} else if rightIdx >= len(right) && leftIdx < len(left) {
			slice[idx] = left[leftIdx]
			leftIdx++
		} else if left[leftIdx] < right[rightIdx] {
			slice[idx] = left[leftIdx]
			leftIdx++
		} else {
			slice[idx] = right[rightIdx]
			rightIdx++
		}
	}
	return slice
}

func MergeSort(slice []int) []int {

	if len(slice) < 2 {
		return slice
	}

	midIdx := (len(slice)) / 2

	return merge(MergeSort(slice[:midIdx]), MergeSort(slice[midIdx:]))
}
