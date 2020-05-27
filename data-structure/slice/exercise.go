// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

// Add imports.
import (
	"fmt"
)

func main() {

	// Declare a nil slice of integers.
	var numbers []int

	// Append numbers to the slice.
	for i := 0; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	// Display each value in the slice.
	for _, value := range numbers {
		fmt.Println(value)
	}

	// Declare a slice of strings and populate the slice with names.
	names := []string{"Madhav", "Keshav", "Raghav", "Mukul"}

	// Display each index position and slice value.
	for i, v := range names {
		fmt.Println("index:", i, "value:", v)
	}

	// Take a slice of index 1 and 2 of the slice of strings.
	newName := names[1:3]

	// Display each index position and slice values for the new slice.
	for i, v := range newName {
		fmt.Println("index:", i, "value:", v)
	}
}
