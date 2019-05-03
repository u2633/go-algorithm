package selection_sort

func SelectionSort(list []int) []int {
	n := len(list)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if list[i] < list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	return list
}
