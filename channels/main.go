package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://toshikiyoshioka.dev",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		time.Sleep(5 * time.Second)
		go checkLink(l, c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " might be down at the moment...")
		c <- link
		return
	}

	fmt.Println(link, " is up and running!")
	c <- link
}