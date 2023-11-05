package main

import "fmt"

func main() {

	// colors := make(map[string]string)
	// colors["green"] = "#4bf745"
	colors := map[string]string {
		"red": "#ff000",
		"white": "#ffffff",
		"green": "#4bf745",
	}

	// delete(colors, "green")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
