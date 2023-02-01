package algorithms_test

import (
	"testing"

	"github.com/telemachus/algorithms"
)

func TestBinarySearch(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		orig   []int
		target int
		want   int
	}{
		"empty slice":                 {[]int{}, 3, -1},
		"wanted not in slice":         {[]int{1, 2, 3}, 4, -1},
		"wanted is a negative number": {[]int{1, 2, 3}, -4, -1},
		"wanted at start of slice":    {[]int{1, 2, 3}, 1, 0},
		"wanted at end of slice":      {[]int{1, 2, 3, 4, 5, 6, 7}, 7, 6},
		"wanted in middle of slice":   {[]int{1, 2, 3, 4, 5, 6, 7}, 4, 3},
	}

	for msg, tc := range tests {
		tc := tc

		t.Run(msg, func(t *testing.T) {
			t.Parallel()

			got := algorithms.BinarySearch(tc.orig, tc.target)
			if got != tc.want {
				t.Errorf(
					"BinarySearch(%v, %v): got %d; want %d",
					tc.orig,
					tc.target,
					got,
					tc.want,
				)
			}
		})
	}
}
