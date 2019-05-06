package sum_medias_convert_time

import "sort"

func SumMediasConvertTime(slice []int) int {
	// sort the slice
	sort.Ints(slice)

	// end calculation when slice length is equal to 2.
	// [1, 2, 3, 4]
	// 3	<- [(1, 2), 3, 4] 	-> append to become a new array [3, 3, 4]
	// 6  	<- [(3, 3), 4]		-> append to become a new array [6, 4]
	// 10 	<- [(6, 4)] 		-> End calculation
	if len(slice) <= 2 {
		return slice[0] + slice[1]
	} else {
		return slice[0] + slice[1] + SumMediasConvertTime(append(slice[2:], slice[0]+slice[1]))
	}
}
