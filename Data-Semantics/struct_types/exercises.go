// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

import "fmt"

// Add imports.

// Add user type and provide comment.
// what is cost of this user struct ?
// Always give preference to simplicity and readablity over performance

type user struct {
	name  string
	email string
	age   int
}

func main() {

	// Declare variable of type user and init using a struct literal.
	user1 := user{
		name:  "Madhav",
		email: "madhav1993hari@gmail.com",
		age:   28,
	}

	// Display the field values.
	fmt.Println("user1", user1.name, user1.age, user1.email)

	// Declare a variable using an anonymous struct.

	user2 := struct {
		name  string
		email string
		age   int
	}{
		name:  "Madhav Jha",
		email: "madhav1994hari@gmail.com",
		age:   28,
	}

	// Display the field values.
	fmt.Println("user2", user2.name, user2.age, user2.email)
}
