package merge_sort

import (
	"fmt"
	"github.com/u2633/go-algorithm/utils"
	"testing"
)

//=== RUN   TestMergeSort
//--- PASS: TestMergeSort (0.02s)
func TestMergeSort(t *testing.T) {
	fmt.Println(MergeSort(utils.GenerateSlice(7)))
}
