package main

import "testing"

func cmp(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range b {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestStringToIntSlice(t *testing.T) {
	results := [][]int{
		[]int{104, 101, 108, 108, 111, 32, 84, 72, 69, 82, 69},
		[]int{49, 52, 47, 37, 51, 55},
		[]int{97, 32, 98, 32, 99, 32, 100, 32, 101, 32, 102, 32, 103, 32, 104, 32, 105, 32, 106, 32, 107, 32, 108, 32, 109, 32, 110, 32, 111, 32, 112, 32, 113, 32, 114, 32, 115, 32, 116, 32, 117, 32, 118, 32, 119, 32, 120, 32, 121, 32, 122},
		[]int{},
	}
	for i := range tests {
		res := StringToIntSlice(tests[i])
		if !cmp(res, results[i]) {
			t.Errorf("Failed with param: '%v'\n"+
				"> %v\n< %v", tests[i], res, results[i])
		}
	}
}
