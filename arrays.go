package main

import "fmt"

func main() {
	urls := [2]string{"google.com", "cotemig.com"}

	fmt.Println(urls)

	for i, url := range urls  {
		urls[i] = "https://" + url

		fmt.Println(url)
	}

	fmt.Println(urls)
}
