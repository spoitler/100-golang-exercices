// Exercise: With a single string variable, use access the first character with the string index
// This is only valid with ASCII characters, and the print value will be the ASCII number
// string[n] is how you should access the value

package main

import "fmt"

func main() {
	var string1 string
	string1 = "hello world"
	fmt.Println(string1[0])
	fmt.Println(string(string1[0]))
}
