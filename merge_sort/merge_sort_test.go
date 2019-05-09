package merge_sort

import (
	"fmt"
	"github.com/u2633/go-algorithm/utils"
	"testing"
)

func TestMergeSort(t *testing.T) {
	// --- PASS: TestMergeSort (0.13s)
	fmt.Println(MergeSort(utils.GenerateSlice(100000)))
	// --- PASS: TestMergeSort (27.37s)
	// MergeSort(utils.GenerateSlice(100000000))
}
