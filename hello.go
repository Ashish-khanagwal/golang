// One more thing about golang is that, In go everything is organized in packages and we're going to use the packages throughout the application.
// The first statement in Go file must be "package..."
// we use "go run <filename>" for the execution.
package main

import "fmt"

// Go program are organized into packages. and Go's standard library provides different core packages for us to use. "fmt" is one of these.

// After it, we have to provide a entrypoint to the Go program, so that it can start exectuing the program. See below
func main() {
	// And here, "main" func is the entry point of a go program
	// at first we will use "go mod init <module path>" This command will create a **go.mod** file.
	// go.mod file describes the "module" with name/module path and go version used in the program.

	fmt.Println("hello world")

	// Variables and Constant
	// In Go, we can declare variables and constants using "var" and "const" keywords.
	var name = "Ashish"
	fmt.Println(name)
	const age = 22
	fmt.Println("My name is", name, "I'm", age, "years old")

	// We can print the data type of the declared variables by using "%T"
	fmt.Printf("The data type of name is %T, data type of age is %T", name, age)

	// Go is statically a typed language
	// Above statement means that, we need to tell the Go compiler, the data type when declaring the variable
	// And, if we are providing the value to the declared variable then Go compiler sets the data type according to the value given.
	var userName string
	// userName = "Ashish"
	var userAge int
	fmt.Println(userName)
	fmt.Println(userAge)

	// We have an alternative to declare the variables in GO language, and this methodology is called "Syntactic Sugar"
	userLocation := "Delhi"
	fmt.Println(userLocation)
	// The above declaration method works with variables not with constants. and also when we explicitly defines data types.
	var locationPincode uint = 124001 // Like here, we can't define the variable as we did above because we defined data type for it.
	fmt.Println(locationPincode)
	// We can also declare multiple variables in a single line
	var (
		userName1 = "Ashish"
		userAge1  = 22
	)
	fmt.Println(userName1, userAge1)

	// Taking Input from the user.
	// We can take input from the user using "fmt.Scanln()" function.
	// basically, Scan(....) saves the user's enterd value in any variable declared in it.
	var playerName string
	fmt.Scan(playerName) // To take the user input. First, we have to understand the concept of Pointer.

	// ########### POINTERS ###########

	// Pointer is a variable that stores the memory address of another variable. And this address points to the actual value of the variable\
	// In Go, we can declare a pointer by using "*" symbol before the variable name.
	// And, we can get the memory address of a variable by using "&" symbol before the variable name.
	var playerAge int = 22
	var playerAgePointer *int = &playerAge
	fmt.Println(playerAgePointer)  // This will print the memory address of the playerAge variable.
	fmt.Println(*playerAgePointer) // This will print the value of the playerAge variable.

}
