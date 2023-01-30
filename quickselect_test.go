package algorithms_test

import (
	"testing"

	"github.com/telemachus/algorithms"
)

var selectTests = map[string]struct {
	given     []int
	nthLowest int
	expected  int
}{
	"empty slice":                           {[]int{}, 0, -1},
	"nthLowest > len(xs)":                   {[]int{1}, 4, -1},
	"one-item slice":                        {[]int{14}, 0, 14},
	"three-item slice; nthLowest at start":  {[]int{15, 16, 14}, 0, 14},
	"three-item slice; nthLowest in middle": {[]int{15, 16, 14}, 1, 15},
	"three-item slice; nthLowest at end":    {[]int{15, 16, 14}, 2, 16},
	"ten-item slice, nthLowest in middle":   {[]int{16, 14, 15, 8, 1, 3, 5, 9, 2, 4}, 4, 5},
}

func TestLomutoQuickselect(t *testing.T) {
	for msg, tc := range selectTests {
		t.Run(msg, func(t *testing.T) {
			actual := algorithms.QuickselectL(tc.given, 0, len(tc.given)-1, tc.nthLowest)

			if tc.expected != actual {
				t.Errorf("expected %d; actual %d", tc.expected, actual)
			}
		})
	}
}

func TestHoareQuickselect(t *testing.T) {
	for msg, tc := range selectTests {
		t.Run(msg, func(t *testing.T) {
			actual := algorithms.QuickselectH(tc.given, 0, len(tc.given)-1, tc.nthLowest)

			if tc.expected != actual {
				t.Errorf("expected %d; actual %d", tc.expected, actual)
			}
		})
	}
}

func TestYourBasicQuickselect(t *testing.T) {
	for msg, tc := range selectTests {
		t.Run(msg, func(t *testing.T) {
			actual := algorithms.QuickselectYB(tc.given, tc.nthLowest)

			if tc.expected != actual {
				t.Errorf("expected %d; actual %d", tc.expected, actual)
			}
		})
	}
}
