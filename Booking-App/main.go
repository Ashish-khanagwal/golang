// One more thing about golang is that, In go everything is organized in packages and we're going to use the packages throughout the application.
// The first statement in Go file must be "package..."
package main

import "fmt"

// Go program are organized into packages. and Go's standard library provides different core packages for us to use. "fmt" is one of these.

// After it, we have to provide a entrypoint to the Go program, so that it can start exectuing the program. See below
func main() {
	// And here, "main" func is the entry point of a go program
	// at first we will use "go mod init <module path>" This command will create a **go.mod** file.
	// go.mod file describes the "module" with name/module path and go version used in the program.
	fmt.Println("hello world")
}

// we use "go run <filename>" for the execution.
