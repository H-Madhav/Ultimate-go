package main

import "fmt"

type user struct {
	name  string
	email string
	age   int64
}

func main() {
	user := user{
		name:  "Madhav",
		email: "madhav1993hari@gmail.com",
		age:   27,
	}

	fmt.Println("Name:", user.name, "Email:", user.email, "age:", user.age)

	user2 := struct {
		name  string
		email string
		age   int64
	}{
		name:  "keshav",
		email: "keshav1993hari@gmail.com",
		age:   27,
	}

	fmt.Println("Name:", user2.name, "Email:", user2.email, "age:", user2.age)

}
