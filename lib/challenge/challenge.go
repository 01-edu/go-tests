package challenge

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"reflect"
	"runtime"
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
			if reflect.TypeOf(v).Kind() == reflect.Func {
				// Function passed as parameter should output the name of the function
				ss[i] = runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name()
			} else {
				// a Go-syntax representation of the value
				ss[i] = fmt.Sprintf("%#v", v)
			}
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
	if err != nil {
		panic(err)
	}
	os.Stdout = w
	out.Results = Call(fn, args)
	outC := make(chan string)
	var buf strings.Builder
	go func() {
		_, _ = io.Copy(&buf, r)
		outC <- buf.String()
	}()
	os.Stdout = old
	_ = w.Close()
	out.Stdout = <-outC
	return out
}

func Function(name string, submittedFunction, expectedFunction interface{}, args ...interface{}) {
	submitted := Monitor(submittedFunction, args)
	expected := Monitor(expectedFunction, args)
	if !reflect.DeepEqual(submitted.Results, expected.Results) {
		Fatalf("%s(%s) == %s instead of %s\n",
			name,
			Format(args...),
			Format(submitted.Results...),
			Format(expected.Results...),
		)
	}
	if !reflect.DeepEqual(submitted.Stdout, expected.Stdout) {
		Fatalf("%s(%s) prints:\n%s\ninstead of:\n%s\n",
			name,
			Format(args...),
			Format(submitted.Stdout),
			Format(expected.Stdout),
		)
	}
}

func Fatal(a ...interface{}) {
	_, _ = fmt.Fprint(os.Stderr, a...)
	os.Exit(1)
}

func Fatalln(a ...interface{}) {
	_, _ = fmt.Fprintln(os.Stderr, a...)
	os.Exit(1)
}

func Fatalf(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
}

func ProgramStdin(exercise, input string, args ...string) {
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
		if input != "" {
			cmd.Stdin = bytes.NewBufferString(input)
		}
		b, err := cmd.CombinedOutput()
		if err != nil {
			if _, ok := err.(*exec.ExitError); !ok {
				return err.Error(), false
			}
			return string(b) + err.Error(), false
		}
		return string(b), true
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
		return fmt.Sprintf(s+"go run . %s\n%s$", strings.Join(quotedArgs, " "), out)
	}
	student, studentOK := run(path.Join("student", exercise))
	solution, solutionOK := run(path.Join("github.com/01-edu/go-tests/solutions", exercise))
	if solutionOK {
		if !studentOK {
			Fatalln("Your program fails (non-zero exit status) when it should not :\n" +
				console(student) +
				"\n\nExpected :\n" +
				console(solution))
		}
	} else {
		if studentOK {
			Fatalln("Your program does not fail when it should (with a non-zero exit status) :\n" +
				console(student) +
				"\n\nExpected :\n" +
				console(solution))
		}
	}
	if student != solution {
		Fatalln("Your program output is not correct :\n" +
			console(student) +
			"\n\nExpected :\n" +
			console(solution))
	}
}

func Program(exercise string, args ...string) {
	ProgramStdin(exercise, "", args...)
}

// TODO: check unhandled errors on all solutions (it should contains "ERROR" on the first line to prove we correctly handle the error)
// TODO: remove the number of rand functions, refactor test cases (aka "table")
