package main

import "fmt"

func main() {
	// var colours map[string]string
	// colours := make(map[string]string)

	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#f123123",
		"white": "#ffffff",
	}
	// colours["white"] = "#ffffff"
	// fmt.Println(colours)
	printMap(colours)
	// delete(colours, "white")
	// fmt.Println(colours)
}

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Println(k, v)
	}
}
