package main

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
)

func workingWithHttpAndFiles() string {

	url := "https://jsonplaceholder.typicode.com/posts"

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("fetching error", err)
	}

	defer response.Body.Close() // cleanup - Close the response body when main() exits
	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("error reading response body", err)
		panic(err)

	}

	fmt.Println("responds nbod", len(responseBody), string(responseBody), reflect.TypeOf((responseBody)))
	return string(responseBody)

}
