package main 

import (
	"net/http"
	"fmt" 
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:%v", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}