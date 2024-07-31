package main

import (
	"log"
	"reflect"
	student "student"
)

var testCases = []struct {
	in   string
	want string
}{
	{
		in:   " ",
		want: "\n",
	},
	{
		in:   "FOR PONY",
		want: "PONY\n",
	},
	{
		in:   "this ... is sparta, then again, maybe not",
		want: "not\n",
	},
	{
		in:   " lorem,ipsum ",
		want: "lorem,ipsum\n",
	},
}

func main() {
	for _, tc := range testCases {
		got := student.LastWord(tc.in)
		if !reflect.DeepEqual(got, tc.want) {
			log.Fatalf("%s(%q) == %q instead of %q\n",
				"LastWord",
				tc.in,
				got,
				tc.want,
			)
		}
	}
}
