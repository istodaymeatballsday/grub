package main

import (
	"fmt"
)

const (
	meat  = iota
	vegan = iota
)

const (
	swedish = iota
	english = iota
)

const (
	// Green color
	Green = "\033[32m"
	// Red color
	Red = "\033[31m"
	// White color
	White = "\033[37m"
)

// PrintColor prints to the terminal with the given ansi color
func PrintColor(color string, s string) {
	fmt.Printf(color + s + "\n")
}

// PrintFood is a wrapper for printColor for the food
func PrintFood(food *Response) {
	printFood((*(food))[meat].DisplayNames[swedish].DishDisplayName)
	printFood((*(food))[vegan].DisplayNames[swedish].DishDisplayName)
}

func printFood(s string) {
	PrintColor(White, "· "+s)
}
