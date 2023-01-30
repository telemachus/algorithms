package algorithms_test

import (
	"testing"

	"github.com/telemachus/algorithms"
)

func TestBinarySearch(t *testing.T) {
	tests := map[string]struct {
		slice    []int
		wanted   int
		expected int
	}{
		"empty slice":                 {[]int{}, 3, -1},
		"wanted not in slice":         {[]int{1, 2, 3}, 4, -1},
		"wanted is a negative number": {[]int{1, 2, 3}, -4, -1},
		"wanted at start of slice":    {[]int{1, 2, 3}, 1, 0},
		"wanted at end of slice":      {[]int{1, 2, 3, 4, 5, 6, 7}, 7, 6},
		"wanted in middle of slice":   {[]int{1, 2, 3, 4, 5, 6, 7}, 4, 3},
	}

	for msg, tc := range tests {
		t.Run(msg, func(t *testing.T) {
			actual := algorithms.BinarySearch(tc.slice, tc.wanted)
			if tc.expected != actual {
				t.Errorf("expected %d; actual %d; given %#v", tc.expected, actual, tc.slice)
			}
		})
	}
}
