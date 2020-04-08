package main

// This is an import block, it allows multiple packages to be imported without
// repeating the "import" keyword
// Unused imports result in compile errors
import (
	"fmt"
)

/*
The main function, when part of the main package, identifies the entry point of
an application.
*/
func main() {
	fmt.Println("Hello, playground")
}
