package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://finvestpaper.ru/",
		"https://fl-bankrotstvo.ru/",
		"https://quiz.fl-bankrotstvo.ru/",
		"https://newquiz.fl-bankrotstvo.ru/",
		"https://golang.org",
	}
	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " might be down!")
		return
	}
	fmt.Println(link + " is up!")
}
