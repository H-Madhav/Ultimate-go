package main

import (
	"fmt"

	"./toy"
)

func main() {

	// Create a value of type toy.
	t := toy.New("Bat", 28)

	// Update the counts.
	t.UpdateOnHand(100)
	t.UpdateSold(2)

	// Display each field separately.
	fmt.Println("Name", t.Name)
	fmt.Println("Weight", t.Weight)
	fmt.Println("OnHand", t.OnHand())
	fmt.Println("Sold", t.Sold())
}
