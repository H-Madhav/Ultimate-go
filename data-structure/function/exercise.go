// Declare a struct type to maintain information about a user. Declare a function
// that creates value of and returns pointers of this type and an error value. Call
// this function from main and display the value.
//
// Make a second call to your function but this time ignore the value and just test
// the error value.
package main

import "fmt"

// user represents a user in the system.
type user struct {
	name  string
	email string
}

// newUser creates and returns pointers of user type values.
func newUser() (*user, error) {
	return &user{"Bill", "bill@ardanlabs.com"}, nil
}

func main() {

	// Create a value of type user.
	u, err := newUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display the value.
	fmt.Println(*u)

	// Call the function and just check the error on the return.
	_, err = newUser()
	if err != nil {
		fmt.Println(err)
		return
	}
}
