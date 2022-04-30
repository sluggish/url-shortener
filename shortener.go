package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	var input string
	fmt.Println("Enter URL: ")
	fmt.Scanln(&input)
	params := url.Values{}
	params.Set(input, "") // encodes user input for get request from api

	request, err := http.Get("https://api.shrtco.de/v2/shorten?url="+params.Encode())
	if err != nil {
		fmt.Println(err)
	}
	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)

	fmt.Println(string(body))

}
