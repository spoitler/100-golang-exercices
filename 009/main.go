// Exercise: Arrays
// Create an array of 10 "int8" values, in it's initialization, fill those values from 0 to 9

package main

func main() {
	ints := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	_ = ints

	// over engineering it with a loop is a loss of energy, and create more CPU operation than creating it manually. 10 it's ok to create it manually
}
