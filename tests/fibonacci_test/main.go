package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"reflect"
	"runtime"

	student "student"
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

// CheckFunctionStructure analyzes if the provided function is recursive and ensures it contains no loops.
func CheckFunctionStructure(f interface{}) error {
	// Get the file path of the loaded function using reflection
	funcPtr := reflect.ValueOf(f).Pointer()
	funcDetails := runtime.FuncForPC(funcPtr)
	funcFile, _ := funcDetails.FileLine(funcPtr)
	funcName := funcDetails.Name() // Retrieve the function name


	fset := token.NewFileSet()

	// Parse the source file to create an AST
	node, err := parser.ParseFile(fset, funcFile, nil, parser.AllErrors)
	if err != nil {
		return fmt.Errorf("failed to parse file: %v", err)
	}

	// Variables to track recursion and loop usage
	var isRecursive, hasLoop bool

	// Traverse the AST to locate the function and check for recursion and loops
	ast.Inspect(node, func(n ast.Node) bool {
		// Locate the function declaration by matching the function name
		if fn, ok := n.(*ast.FuncDecl); ok && fn.Name.Name == "Fibonacci" {

			// Traverse the function body to check for recursive calls and loops
			ast.Inspect(fn.Body, func(n ast.Node) bool {
				switch stmt := n.(type) {
				case *ast.CallExpr:
					// Check if the function calls itself (recursion)
					if ident, ok := stmt.Fun.(*ast.Ident); ok && ident.Name == fn.Name.Name {
						isRecursive = true // Recursive call detected
					}
				case *ast.ForStmt:
					// Detect any "for" loop
					hasLoop = true
				case *ast.RangeStmt:
					// Detect any "range" loop
					hasLoop = true
				}
				return true
			})
			return false // Stop further inspection after analyzing the function
		}
		return true
	})

	// Check results and return appropriate errors
	if hasLoop {
		return fmt.Errorf("the function %s contains a loop, which is not allowed", funcName)
	}
	if !isRecursive {
		return fmt.Errorf("the function %s is not recursive", funcName)
	}
	return nil
}

func main() {
	// Check the structure of the student's Fibonacci implementation
	err := CheckFunctionStructure(student.Fibonacci)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	// Generate test cases for the function
	args := append(random.IntSliceBetween(0, 25),
		4,
		5,
		-5,
	)
	for _, arg := range args {
		challenge.Function("Fibonacci", student.Fibonacci, solutions.Fibonacci, arg)
	}
}
