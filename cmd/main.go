package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println("file system examples", fileSystem())
	x, y := Split(12)
	fmt.Println(":x is", x, "y is", y)
	// fmt.Println("example app is", randomNumber(), workingWithHttpAndFiles())

}
