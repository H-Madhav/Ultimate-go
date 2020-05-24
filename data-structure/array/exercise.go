// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
// Display the string value and address of each element.
package main

import "fmt"

// Add imports.

func main() {

	// Declare an array of 5 strings set to its zero value.
	var arr1 [5]string

	// Declare an array of 5 strings and pre-populate it with names.
	arr2 := [5]string{"Hello", "there", "How", "are", "You"}

	// Assign the populated array to the array of zero values.
	arr1 = arr2

	// Iterate over the first array declared.
	// Display the string value and address of each element.
	for i, v := range arr1 {
		fmt.Println("value:", v, "address:", &arr1[i], "Address of V:", &v)
		// how the behavior of the for range and
		// how memory for an array is contiguous.
	}
}
