package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func main() {
	tests := []struct {
		arr       string
		targetSum string
		want      string
	}{
		{
			arr:       "[1, 2, 3, 4, 5]",
			targetSum: "6",
			want:      "Pairs with sum 6: [[0 4] [1 3]]\n",
		},
		{
			arr:       "[1, 2, 3, 4, 5, 1]",
			targetSum: "6",
			want:      "Pairs with sum 6: [[0 4] [1 3] [4 5]]\n",
		},
		{
			arr:       "[-1, 2, -3, 4, -5]",
			targetSum: "1",
			want:      "Pairs with sum 1: [[0 1] [2 3]]\n",
		},
		{
			arr:       "[-1, -2, -3, -4, -5]",
			targetSum: "-5",
			want:      "Pairs with sum -5: [[0 3] [1 2]]\n",
		},
		{
			arr:       "[1, 2, 3, 4, 5]",
			targetSum: "10",
			want:      "No pairs found.\n",
		},
		{
			arr:       "[1, 2, 3, 4, 20, -4, 5]",
			targetSum: "2 5",
			want:      "Invalid target sum.\n",
		},
		{
			arr:       "[1, 2, 3, 4, 20, -4, 5]",
			targetSum: "l",
			want:      "Invalid target sum.\n",
		},
		{
			arr:       "[1, 2, 3, 4, 20, p, 5]",
			targetSum: "5",
			want:      "Invalid number: p\n",
		},
		{
			arr:       "[1, 2, 3, 4, 20,    p, 5]",
			targetSum: "5",
			want:      "Invalid number: p\n",
		},
		{
			arr:       "[1, 2, 3, 4, 20, 5",
			targetSum: "5",
			want:      "Invalid input.\n",
		},
		{
			arr:       "1, 2, 3, 4, 20, 5",
			targetSum: "5",
			want:      "Invalid input.\n",
		},
		{
			arr:       "1 2 3 4 20 5",
			targetSum: "5",
			want:      "Invalid input.\n",
		},
	}

	for _, tc := range tests {
		got, _ := runStudentProgram("findpairs", []string{tc.arr, tc.targetSum}...)
		if got != tc.want {
			fmt.Printf("findpairs %q %q -> ", tc.arr, tc.targetSum)
			fmt.Printf("got: %q instead of %q\n", got, tc.want)
			os.Exit(1)
		}
	}
}

func runStudentProgram(exercise string, args ...string) (string, bool) {
	run := func(pkg string) (string, bool) {
		binaryPath := path.Join(os.TempDir(), "binaries", path.Base(path.Dir(pkg)), path.Base(pkg))
		if _, err := os.Stat(binaryPath); os.IsNotExist(err) {
			if b, err := exec.Command("go", "build", "-o", binaryPath, pkg).CombinedOutput(); err != nil {
				return string(b), false
			}
		}
		if fi, err := os.Stat(binaryPath); err != nil || fi.Mode()&0111 == 0 {
			return "go run: cannot run non-main package\n", false
		}
		cmd := exec.Command(binaryPath, args...)
		b, err := cmd.CombinedOutput()
		if err != nil {
			if _, ok := err.(*exec.ExitError); !ok {
				return err.Error(), false
			}
			return string(b) + err.Error(), false
		}
		return string(b), true
	}
	return run(path.Join("student", exercise))
}
