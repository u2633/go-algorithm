package bubble_sort

func BubbleSort(list []int) []int {
	n := len(list)

	for n > 1 {
		newN := 0
		for i := 1; i <= n-1; i++ {
			if list[i-1] > list[i] {
				list[i-1], list[i] = list[i], list[i-1]
				newN = i
			}
		}
		n = newN
	}
	return list
}

func OldBubbleSort(list []int) []int {
	n := len(list)

	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}
