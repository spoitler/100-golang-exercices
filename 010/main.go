// Create an array of 5 strings, and initialize it's 2 first values with some random names

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	names := [5]string{}
	names[0] = randomName()
	names[1] = randomName()
	names[2] = "alice"
	names[3] = "bob"
	names[4] = "claire"
	for _, name := range names {
		fmt.Println(name)
	}
}

func randomName() string {
	var name string
	for i := 0; i < rand.Intn(10); i++ {
		name += string(rune(rand.Intn(25) + 97))
	}
	return name
}
