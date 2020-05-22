// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

import "fmt"

// Add imports

// main is the entry point for the application.
func main() {

	// Declare variables that are set to their zero value.
	var a int
	var b string
	var c bool

	// Display the value of those variables.

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)

	// Declare variables and initialize.
	// Using the short variable declaration operator.

	aa := 1
	bb := "Hello"
	cc := true

	// Display the value of those variables.
	fmt.Println("aa:", aa)
	fmt.Println("bb:", bb)
	fmt.Println("cc:", cc)

	// Perform a type conversion.
	pi := float32(3.14)

	// Display the value of that variable.
	fmt.Printf("%T [%v]\n", pi, pi)
}
