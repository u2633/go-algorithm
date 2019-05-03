package selection_sort

import (
	"math/rand"
	"testing"
	"time"
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	l = 100000
)

/*
=== RUN   TestSelectionSort
--- PASS: TestSelectionSort (14.84s)
*/
func TestSelectionSort(t *testing.T) {
	SelectionSort(r.Perm(l))
}
