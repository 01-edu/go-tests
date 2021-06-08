package main

import (
	"reflect"
	"runtime"
	"sort"
	"strings"

	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func advancedSortWordArr(a []string, f func(a, b string) int) {
	sort.Slice(a, func(i, j int) bool {
		return f(a[i], a[j]) < 0
	})
}

func main() {
	args := [][]string{{"a", "A", "1", "b", "B", "2", "c", "C", "3"}}

	args = append(args, random.StrSlice(chars.Words))

	for _, arg := range args {
		// copy for using the solution function
		cp_sol := make([]string, len(arg))
		// copy for using the student function
		cp_stu := make([]string, len(arg))

		copy(cp_sol, arg)
		copy(cp_stu, arg)

		advancedSortWordArr(cp_sol, strings.Compare)
		student.AdvancedSortWordArr(cp_stu, strings.Compare)

		if !reflect.DeepEqual(cp_stu, cp_sol) {
			challenge.Fatalf("%s(%v, %s) == %v instead of %v\n",
				"AdvancedSortWordArr",
				arg,
				runtime.FuncForPC(reflect.ValueOf(strings.Compare).Pointer()).Name(),
				cp_stu,
				cp_sol,
			)
		}
	}
}
