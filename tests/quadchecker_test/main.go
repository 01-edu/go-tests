package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	execFatal := func(name string, arg ...string) string {
		b, err := exec.Command(name, arg...).CombinedOutput()
		if err != nil {
			challenge.Fatalf("%s\n", err)
		}
		return string(b)
	}

	// only one output no matter the arguments
	// output :
	// quadA
	// quadB
	quadAB := []string{
		"o",
		"/",
	}

	// special cases one output
	specialCases := []string{
		// quadC
		"AA",
		// quadD
		"A\nA",
	}

	multipleOutputs := []string{
		// tree output : quadC quadD QuadE
		"A",
		// two output : quadC quadE
		"A\nC",
		// quadD quadE
		"AC",
	}

	solBinary := "quadcheckerprog_solution"
	studBinary := "quadcheckerprog_student"

	defer removeBinary(solBinary, studBinary)
	execFatal("go", "build", "-o", solBinary, path.Join("github.com/01-edu/go-tests/solutions", "quadchecker"))
	execFatal("go", "build", "-o", studBinary, path.Join("student", "quadchecker"))

	testCases := [][]string{quadAB, specialCases, multipleOutputs}
	for _, c := range testCases {
		for _, s := range c {
			cmdCorrect := fmt.Sprintf("echo -e '%s' | ./%s ", s, solBinary)
			cmdStudent := fmt.Sprintf("echo -e '%s' | ./%s", s, studBinary)
			correct := execFatal("bash", "-c", cmdCorrect)
			output := execFatal("bash", "-c", cmdStudent)
			if output != correct {
				removeBinary(solBinary, studBinary)
				challenge.Fatalf("./quadchecker prints %q instead of %q\n", output, correct)
			}
		}
	}
}

func removeBinary(s ...string) {
	for _, c := range s {
		e := os.Remove(c)
		if e != nil {
			challenge.Fatal(e)
		}
	}
}
