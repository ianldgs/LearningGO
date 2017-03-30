package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func getHTML(url string) (result string, err error) {
	response, err := http.Get(url)

	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
		body = make([]byte, 0);
	}

	result = string(body)

	return
}

func main() {
	getHTML("https://www.cotemig.com.br")
}