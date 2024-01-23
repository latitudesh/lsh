package utils

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
)

func BodyReader(response runtime.ClientResponse) {
	defer response.Body().Close()

	// Read the response body into a byte slice
	bodyBytes, err := io.ReadAll(response.Body())
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	// Convert the byte slice to a string
	bodyString := string(bodyBytes)

	// Print the string
	fmt.Println(bodyString)
}
