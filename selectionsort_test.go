package algorithms_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/telemachus/algorithms"
)

func TestSelectionSort(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		orig []int
		want []int
	}{
		"empty slice":                 {[]int{}, []int{}},
		"one-item slice":              {[]int{14}, []int{14}},
		"three-item slice":            {[]int{15, 16, 14}, []int{14, 15, 16}},
		"four-item slice":             {[]int{15, 16, 17, 14}, []int{14, 15, 16, 17}},
		"five-item slice":             {[]int{5, 6, 7, 4, 8}, []int{4, 5, 6, 7, 8}},
		"slice containing duplicates": {[]int{2, 4, 2, 3, 1, 2}, []int{1, 2, 2, 2, 3, 4}},
		"ten-item slice": {
			[]int{16, 14, 15, 8, 1, 3, 5, 9, 2, 4},
			[]int{1, 2, 3, 4, 5, 8, 9, 14, 15, 16},
		},
		"ten-item slice, reversed": {
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for msg, tc := range tests {
		tc := tc

		t.Run(msg, func(t *testing.T) {
			t.Parallel()

			got := make([]int, len(tc.orig))
			copy(got, tc.orig)
			algorithms.SelectionSort(got)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf(
					"SelectionSort(%#v) failure (-want +got)\n%s",
					tc.orig,
					diff,
				)
			}
		})
	}
}
