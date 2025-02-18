// Exercise: User input
// Using only the fmt package, ask a user for it's name and then for it's surname
// Store it in 2 variables called "name" and "surname"
// After user has entered the data, print it out

// Tip: https://pkg.go.dev/fmt#hdr-Scanning

package main

import "fmt"

func main() {
	var name, surname string
	fmt.Println("What's your name and surname ? enter it separated by a space")
	_, err := fmt.Scan(&name, &surname)
	if err != nil {
		panic(err)
	}
	fmt.Printf("you'r name is %v and surname is %v\n", name, surname)
}
