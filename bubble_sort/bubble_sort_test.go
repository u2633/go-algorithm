package bubble_sort

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
=== RUN   TestBubbleSort
--- PASS: TestBubbleSort (15.87s)
*/
func TestBubbleSort(*testing.T) {
	BubbleSort(r.Perm(l))
}

/*
=== RUN   TestOldBubbleSort
--- PASS: TestOldBubbleSort (18.68s)
*/
func TestOldBubbleSort(t *testing.T) {
	OldBubbleSort(r.Perm(l))
}
