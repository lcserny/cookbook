package main

import (
	"cookbook/chapter2/ansicolor"
	"fmt"
)

func main() {
	c := ansicolor.ColorText{
		TextColor: ansicolor.Red,
		Text:      "I'm red!",
	}
	fmt.Println(c.String())

	c.TextColor = ansicolor.Green
	c.Text = "Now I'm green!"
	fmt.Println(c.String())

	c.TextColor = ansicolor.ColorNone
	c.Text = "Back to normal..."
	fmt.Println(c.String())
}
