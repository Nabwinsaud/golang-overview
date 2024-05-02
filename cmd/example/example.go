package example

import (
	"fmt"
	"strconv"
)

func ExampleFunc() int {
	str1 := "10"
	num, err := strconv.Atoi((str1))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	return num
}
