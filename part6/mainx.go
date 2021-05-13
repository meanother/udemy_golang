package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://finvestpaper.ru")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	// fmt.Println(resp)
	// fmt.Println(resp.Status)
	// fmt.Println(resp.StatusCode)
	// fmt.Println(resp.Header)
	// fmt.Println(resp.Body)

	// // bs := []byte{}
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}

1-1
2-2
3-1
4-3
5-3
6-1
7-1
8-1