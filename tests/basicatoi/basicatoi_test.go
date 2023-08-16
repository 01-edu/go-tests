package main

import "testing"

func TestBasicAtoi(t *testing.T) {
	for i := range tests {
		res := BasicAtoi(tests[i])
		if res != results[i] {
			t.Errorf("Failed with param: %v\n"+
				"> %v\n"+
				"< %v\n",
				tests[i],
				results[i],
				res,
			)
		}
	}
}
