package main

import (
	"log"
	"reflect"
	"strconv"
	"strings"
	student "student"
)

func main() {
	tests := []struct {
		args []uint
		want bool
	}{
		{args: []uint{2, 3, 1, 1, 4}, want: true},
		{args: []uint{1, 1, 1, 1, 0}, want: true},
		{args: []uint{5, 4, 3, 2, 1, 0}, want: true},
		{args: []uint{0}, want: true},
		{args: []uint{5}, want: true},
		{args: []uint{}, want: false},
		{args: []uint{1, 2, 3, 0, 2}, want: false},
		{args: []uint{3, 2, 1, 0, 4}, want: false},
		{args: []uint{0, 0, 0, 0, 0}, want: false},
		{args: []uint{1, 2, 3}, want: false},
		{args: []uint{1, 2, 3, 0, 1}, want: false},
		{args: []uint{1, 0, 0, 0, 0}, want: false},
		{args: []uint{1, 0, 1, 0, 1}, want: false},
		{args: []uint{10, 20, 30, 40, 0}, want: false},
	}

	for _, tc := range tests {
		got := student.CanJump(tc.args)
		if !reflect.DeepEqual(got, tc.want) {
			log.Fatalf("%s(%#v) == %#v instead of %#v\n",
				"CanJump",
				uintSliceToString(tc.args),
				got,
				tc.want,
			)
		}
	}
}

func uintSliceToString(nums []uint) string {
	var strSlice []string
	for _, num := range nums {
		strSlice = append(strSlice, strconv.FormatUint(uint64(num), 10))
	}
	return "[" + strings.Join(strSlice, ", ") + "]"
}
