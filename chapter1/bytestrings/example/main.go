package main

import (
	"cookbook/chapter1/bytestrings"
	"fmt"
)

func main() {
	err := bytestrings.WorkWithBuffer()
	if err != nil {
		panic(err)
	}
	fmt.Println()

	bytestrings.SearchString()
	bytestrings.ModifyString()
	bytestrings.StringReader()

	fmt.Println()
}
