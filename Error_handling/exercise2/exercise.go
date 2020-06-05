// Create a custom error type called appError that contains three fields, err error,
// message string and code int. Implement the error interface providing your own message
// using these three fields. Implement a second method named temporary that returns false
// when the value of the code field is 9. Write a function called checkFlag that accepts
// a bool value. If the value is false, return a pointer of your custom error type
// initialized as you like. If the value is true, return a default error. Write a main
// function to call the checkFlag function and check the error using the temporary
// interface.
package main

import (
	"errors"
	"fmt"
)

// Add imports.

// Declare a struct type named appError with three fields, err of type error,
// message of type string and code of type int.
type appError struct {
	err     error
	message string
	code    int
}

// Declare a method for the appError struct type that implements the
// error interface.

func (a *appError) Error() string {
	return fmt.Sprintf("App Error[%s] Message[%s] Code[%d]", a.err, a.message, a.code)
}

// Declare a method for the appError struct type named Temporary that returns
// true when the value of the code field is not 9.

func (a *appError) Temporary() bool {
	return (a.code != 9)
}

// Declare the temporary interface type with a method named Temporary that
// takes no parameters and returns a bool.

type temporary interface {
	Temporary() bool
}

// Declare a function named checkFlag that accepts a boolean value and
// returns an error interface value.

func checkFlag(t bool) error {

	if !t {
		return &appError{errors.New("Flag False"), "The Flag was false", 9}
	}

	// Return a default error.
	return errors.New("Flag True")
}

func main() {

	if err := checkFlag(false); err != nil {
		switch e := err.(type) {
		case temporary:
			fmt.Println(err)
			if !e.Temporary() {
				fmt.Println("Critical Error!")
			}
		default:
			fmt.Println(err)
		}
	}
}
