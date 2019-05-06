package sum_medias_convert_time

import (
	"fmt"
	"testing"
)

func TestSumMediasConvertTime(t *testing.T) {
	medias1 := []int{8, 6, 4, 12}
	fmt.Println(SumMediasConvertTime(medias1))

	medias2 := []int{1, 2, 3, 4}
	fmt.Println(SumMediasConvertTime(medias2))

	medias3 := []int{9, 8, 7, 6}
	fmt.Println(SumMediasConvertTime(medias3))

	medias4 := []int{10, 1, 7, 4}
	fmt.Println(SumMediasConvertTime(medias4))
}
