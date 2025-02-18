// Exercise: Conditional
// Check if the range of a number is between 20 and 30
// If the number is below 20 print : too cold
// If the number is inbetween print: perfect
// If the number is above 30 print : so hot

// Use if and a else if!

package main

import "fmt"

func main() {
	fmt.Println("Give me a number")
	var number int
	_, err := fmt.Scanln(&number)
	if err != nil {
		panic(err)
	}
	// Here goes your code
	if number > 30 {
		fmt.Println("so hot")
	} else if number < 30 && number > 20 {
		fmt.Println("perfect")
	} else {
		fmt.Println("so cold")
	}
}
