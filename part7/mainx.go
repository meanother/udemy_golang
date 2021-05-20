package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://finvestpaper.ru/",
		"https://fl-bankrotstvo.ru/",
		"https://quiz.fl-bankrotstvo.ru/",
		"https://newquiz.fl-bankrotstvo.ru/",
		"https://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }
	// for {
	// 	go checkLink(<-c, c)
	// }

	for l := range c {
		go func(link string) {
			time.Sleep(time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	// time.Sleep(time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " might be down!")
		c <- link
		return
	}
	fmt.Println(link + " is up!")
	// c <- "Yes, its up!"
	c <- link
}


1-2
2-3
3-4
4-2
5-1
6-1