package main

import (
	"cookbook/chapter2/flags"
	"flag"
	"fmt"
)

func main() {
	// initialize our Setup
	c := flags.Config{}
	c.Setup()

	// generally call this from main
	flag.Parse()

	fmt.Println(c.GetMessage())
}
