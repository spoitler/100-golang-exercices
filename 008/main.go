// Exercise: User input
// Get a number from the console and check if it's odd
// You can create another function or do it everything in the main func :)

package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter a number and i tell you if it's odd or not")
	_, err := fmt.Scan(&number)
	if err != nil {
		panic(err)
	}
	isTheNUmberOdd := isOdd(number)
	if isTheNUmberOdd {
		fmt.Printf("you number is odd")
		return
	}
	fmt.Printf("you number is even")
}

func isOdd(n int) bool {
	if n%2 == 0 {
		return false
	}
	return true
}
