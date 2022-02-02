package main

import (
	"fmt"
	"log"
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

	defer rem("quadcheckerprog_solution", "quadcheckerprog_student")
	execFatal("go", "build", "-o", "quadcheckerprog_solution", "../../solutions/quadchecker")
	execFatal("go", "build", "-o", "quadcheckerprog_student", path.Join("student", "quadchecker"))

	testCases := [][]string{quadAB, specialCases, multipleOutputs}
	for _, c := range testCases {
		for _, s := range c {
			cmdCorrect := fmt.Sprintf("echo -e '%s' | ./quadcheckerprog_solution", s)
			cmdStudent := fmt.Sprintf("echo -e '%s' | ./quadcheckerprog_student", s)
			correct := execFatal("bash", "-c", cmdCorrect)
			output := execFatal("bash", "-c", cmdStudent)
			if output != correct {
				challenge.Fatalf("./quadchecker prints %q instead of %q\n", output, correct)
			}
		}
	}
}

func rem(s ...string) {
	for _, c := range s {
		e := os.Remove(c)
		if e != nil {
			log.Fatal(e)
		}
	}
}
