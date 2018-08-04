package ansicolor

import "fmt"

// color of text
type Color int

const (
	// ColorNone is default
	ColorNone = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	Black Color = -1
)

// ColorText holds a string and its color
type ColorText struct {
	TextColor Color
	Text      string
}

func (c *ColorText) String() string {
	if c.TextColor == ColorNone {
		return c.Text
	}

	value := 30
	if c.TextColor != Black {
		value += int(c.TextColor)
	}
	// \033[0;31mtest\033[0m
	return fmt.Sprintf("\033[0;%dm%s\033[0m", value, c.Text)
}
