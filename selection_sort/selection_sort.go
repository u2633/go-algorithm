package selection_sort

func SelectionSort(list []int) []int {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] < list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	return list
}
