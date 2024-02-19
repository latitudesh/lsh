package utils

import (
	"fmt"
	"os"
)

func Exit(msg string, err error) {
	fmt.Printf("%v: %v\n", msg, err)
	os.Exit(1)
}
