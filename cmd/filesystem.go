package main

import (
	"fmt"
	"os"
)

func fileSystem() []string {
	args := os.Args

	// fmt.Println("args ", args)

	homeDir := os.Getenv("HOME")

	data := []byte("hello world\n let code in golang where concurrency is king\n")
	err := os.WriteFile("example.txt", data, 0644)
	if err != nil {
		fmt.Println("error writing to file", err)
	}

	fmt.Println("file written successfully")
	content, err := os.ReadFile("./example.txt")

	if err != nil {
		fmt.Println("error reading file", err)
	}
	fmt.Println("file content is", string(content), content)
	fmt.Println("home dir is", homeDir)
	for i := 0; i < len(args); i++ {
		fmt.Println("arg is", args[i])
	}

	i := 0

	for {
		fmt.Println("iteration is ", i)
		if i == 100 {
			break
		}
		i++
	}

	entry := []string{"Nabin", "khem", "Ram", "Hari", "Shyam"}

	for i, name := range entry {
		fmt.Printf("index sis %d and name is %s\n", i, name)
	}
	return args
}
