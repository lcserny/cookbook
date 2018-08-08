package main

import (
	"cookbook/chapter3/currency"
	"fmt"
)

func main() {
	// start with our user input
	// of fifteen dollars and 93 cents
	userInput := "15.93"

	pennies, err := currency.ConvertStringDollarsToPennies(userInput)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User input converted to %d pennies\n", pennies)

	// adding 15 cents
	pennies += 15

	dollars := currency.ConvertPenniesToDollarString(pennies)
	fmt.Printf("Added 15 cents, new value is %d dollars\n", dollars)
}
