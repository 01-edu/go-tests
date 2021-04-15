package challenge

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"reflect"
	"strconv"
	"strings"
)

func Format(a ...interface{}) string {
	ss := make([]string, len(a))
	for i, v := range a {
		switch v.(type) {
		case nil:
			ss[i] = "nil" // instead of "<nil>"
		case
			string,
			byte, // uint8
			rune: // int32

			// string     : a double-quoted string safely escaped with Go syntax
			// byte, rune : a single-quoted character literal safely escaped with Go syntax
			ss[i] = fmt.Sprintf("%q", v)
		default:
			// a Go-syntax representation of the value
			ss[i] = fmt.Sprintf("%#v", v)
		}
	}
	return strings.Join(ss, ", ")
}

func Call(fn interface{}, args []interface{}) []interface{} {
	// Convert args from []interface{} to []reflect.Value
	vals := make([]reflect.Value, len(args))
	for i, v := range args {
		if v != nil {
			vals[i] = reflect.ValueOf(v)
		} else {
			vals[i] = reflect.Zero(reflect.TypeOf((*interface{})(nil)).Elem())
		}
	}

	vals = reflect.ValueOf(fn).Call(vals)

	// Convert the return values from []reflect.Value to []interface{}
	result := make([]interface{}, len(vals))
	for i, v := range vals {
		result[i] = v.Interface()
	}
	return result
}

type Output struct {
	Results []interface{}
	Stdout  string
}

func Monitor(fn interface{}, args []interface{}) (out Output) {
	old := os.Stdout
	r, w, err := os.Pipe()
	panicIfNotNil(err)
	os.Stdout = w
	out.Results = Call(fn, args)
	outC := make(chan string)
	var buf strings.Builder
	go func() {
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	os.Stdout = old
	w.Close()
	out.Stdout = <-outC
	return out
}

func Function(name string, fn1, fn2 interface{}, args ...interface{}) {
	st1 := Monitor(fn1, args)
	st2 := Monitor(fn2, args)
	if !reflect.DeepEqual(st1.Results, st2.Results) {
		Fatalf("%s(%s) == %s instead of %s\n",
			name,
			Format(args...),
			Format(st1.Results...),
			Format(st2.Results...),
		)
	}
	if !reflect.DeepEqual(st1.Stdout, st2.Stdout) {
		Fatalf("%s(%s) prints:\n%s\ninstead of:\n%s\n",
			name,
			Format(args...),
			Format(st1.Stdout),
			Format(st2.Stdout),
		)
	}
}

func Fatal(a ...interface{}) {
	fmt.Fprint(os.Stderr, a...)
	os.Exit(1)
}

func Fatalln(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
	os.Exit(1)
}

func Fatalf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
}

func panicIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}

func execute(cmd *exec.Cmd) (string, int) {
	b, err := cmd.CombinedOutput()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return string(b), ee.ExitCode()
		}
		panic(err)
	}
	return string(b), 0
}

func ProgramStdin(exercise, input string, args ...string) {
	run := func(pkg string) (string, int) {
		binaryPath := path.Join(os.TempDir(), "binaries", path.Base(path.Dir(pkg)), path.Base(pkg))
		if s, code := execute(exec.Command("go", "build", "-o", binaryPath, pkg)); code != 0 {
			return s, code
		}
		if _, err := os.Stat(binaryPath); os.IsNotExist(err) {
			return "Failed to compile your code as a program", 1
		}
		cmd := exec.Command(binaryPath, args...)
		if input != "" {
			cmd.Stdin = bytes.NewBufferString(input)
		}
		return execute(cmd)
	}
	console := func(out string) string {
		var quotedArgs []string
		for _, arg := range args {
			quotedArgs = append(quotedArgs, strconv.Quote(arg))
		}
		s := "\n$ "
		if input != "" {
			s += "echo -ne " + strconv.Quote(input) + " | "
		}
		return fmt.Sprintf(s+"./%s %s\n%s$ ", exercise, strings.Join(quotedArgs, " "), out)
	}
	code := func(code int) string {
		return fmt.Sprintf("echo $?\n%d\n$", code)
	}
	student, studentCode := run(path.Join("student", exercise))
	solution, solutionCode := run(path.Join("github.com/01-edu/go-tests/solutions", exercise))
	if solutionCode == 0 {
		if studentCode != 0 {
			Fatalln("Your program fails (non-zero exit status) when it should not :\n" +
				console(student) +
				code(studentCode) + "\n\n" +
				"Expected :\n" +
				console(solution) +
				code(solutionCode))
		}
	} else {
		if studentCode == 0 {
			Fatalln("Your program does not fail when it should (with a non-zero exit status) :" + "\n" +
				console(student) +
				code(studentCode) + "\n\n" +
				"Expected :\n" +
				console(solution) +
				code(solutionCode))
		}
	}
	if student != solution {
		Fatalln("Your program output is not correct :\n" +
			console(student) + "\n\n" +
			"Expected :\n" +
			console(solution))
	}
}

func Program(exercise string, args ...string) {
	ProgramStdin(exercise, "", args...)
}

// TODO: check unhandled errors on all solutions (it should contains "ERROR" on the first line to prove we correctly handle the error)
// TODO: remove the number of rand functions, refactor test cases (aka "table")
