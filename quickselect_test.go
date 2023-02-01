package algorithms_test

import (
	"testing"

	"github.com/telemachus/algorithms"
)

var selectTests = map[string]struct {
	orig      []int
	nthLowest int
	want      int
}{
	"empty slice":                              {[]int{}, 0, -1},
	"nthLowest > len(xs)":                      {[]int{1}, 4, -1},
	"one-item slice":                           {[]int{14}, 0, 14},
	"three-item slice; nthLowest at start":     {[]int{14, 16, 15}, 0, 14},
	"three-item slice; nthLowest in middle":    {[]int{16, 15, 14}, 1, 15},
	"three-item slice; nthLowest at end":       {[]int{15, 14, 16}, 2, 16},
	"three-item slice; nthLowest is duplicate": {[]int{14, 14, 16}, 1, 14},
	"ten-item slice, nthLowest in middle":      {[]int{16, 14, 15, 8, 1, 3, 5, 9, 2, 4}, 4, 5},
}

func TestLomutoQuickselect(t *testing.T) {
	t.Parallel()

	for msg, tc := range selectTests {
		tc := tc

		t.Run(msg, func(t *testing.T) {
			t.Parallel()

			dupe := make([]int, len(tc.orig))
			copy(dupe, tc.orig)
			got := algorithms.QuickselectL(
				dupe,
				0,
				len(dupe)-1,
				tc.nthLowest,
			)

			if got != tc.want {
				t.Errorf(
					"QuickselectL(%#v) == %d; want %d",
					tc.orig,
					got,
					tc.want,
				)
			}
		})
	}
}

func TestHoareQuickselect(t *testing.T) {
	t.Parallel()

	for msg, tc := range selectTests {
		tc := tc

		t.Run(msg, func(t *testing.T) {
			t.Parallel()

			dupe := make([]int, len(tc.orig))
			copy(dupe, tc.orig)
			got := algorithms.QuickselectH(
				dupe,
				0,
				len(dupe)-1,
				tc.nthLowest,
			)

			if got != tc.want {
				t.Errorf(
					"QuickselectH(%#v) == %d; want %d",
					tc.orig,
					got,
					tc.want,
				)
			}
		})
	}
}

func TestYourBasicQuickselect(t *testing.T) {
	t.Parallel()

	for msg, tc := range selectTests {
		tc := tc

		t.Run(msg, func(t *testing.T) {
			t.Parallel()

			dupe := make([]int, len(tc.orig))
			copy(dupe, tc.orig)
			got := algorithms.QuickselectYB(dupe, tc.nthLowest)

			if got != tc.want {
				t.Errorf(
					"QuickselectYB(%#v) == %d; want %d",
					tc.orig,
					got,
					tc.want,
				)
			}
		})
	}
}
