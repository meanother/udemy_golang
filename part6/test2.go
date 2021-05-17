package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := readArgs()
	fmt.Println(fileName)
	readFile(fileName)
}

func readArgs() string {
	a := os.Args
	return a[1]
}

func readFile(a string) {
	file, err := os.Open(a)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	// fmt.Println(file)
	io.Copy(os.Stdout, file)
}
