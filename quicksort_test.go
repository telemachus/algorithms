package algorithms_test

import (
	"reflect"
	"testing"

	"github.com/telemachus/algorithms"
)

var tests = map[string]struct {
	given    []int
	expected []int
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

func TestQuicksortL(t *testing.T) {
	for msg, tc := range tests {
		t.Run(msg, func(t *testing.T) {
			algorithms.QuicksortL(tc.given, 0, len(tc.given)-1)

			if !reflect.DeepEqual(tc.expected, tc.given) {
				t.Errorf("expected %#v; actual %#v", tc.expected, tc.given)
			}
		})
	}
}

func TestQuicksortH(t *testing.T) {
	for msg, tc := range tests {
		t.Run(msg, func(t *testing.T) {
			algorithms.QuicksortH(tc.given, 0, len(tc.given)-1)

			if !reflect.DeepEqual(tc.expected, tc.given) {
				t.Errorf("expected %#v; actual %#v", tc.expected, tc.given)
			}
		})
	}
}

func TestQuicksortYB(t *testing.T) {
	for msg, tc := range tests {
		t.Run(msg, func(t *testing.T) {
			algorithms.QuicksortYB(tc.given)

			if !reflect.DeepEqual(tc.expected, tc.given) {
				t.Errorf("expected %#v; actual %#v", tc.expected, tc.given)
			}
		})
	}
}
