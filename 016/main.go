// Exercise: Check if a file exists
// Tip: use the "os" package

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("test.txt"); errors.Is(err, os.ErrNotExist) {
		fmt.Println("test.txt does not exist")
	}
}
