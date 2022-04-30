package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var input string
	fmt.Println("Enter URL:")
	fmt.Scanln(&input)

	request, err := http.Get("https://api.shrtco.de/v2/shorten?url="+input)
	if err != nil {
		fmt.Println(err)
	}
	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)

	fmt.Println(string(body))

}
