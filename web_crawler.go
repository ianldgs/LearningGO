package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"regexp"
)

var regex_links = regexp.MustCompile("(?U)href=\"(.*)\"")
var regex_dont_visit = regexp.MustCompile("(\\.js|\\.css|\\.ico)$")
var regex_uri = regexp.MustCompile(`.*\..*`)

func FindUrls(body string) (urls []string) {
	for _, url := range regex_links.FindAllStringSubmatch(body, -1) {
		urls = append(urls, url[1])
	}

	return
}

func findUrls(body string, c chan string) {
	for _, url := range FindUrls(body) {
		c <- url
	}
}

func Crawl(url string, depth int) (urls []string) {
	response, _ := http.Get(url)

	body, _ := ioutil.ReadAll(response.Body)

	urls = FindUrls(string(body))

	return
}

var visited = make(map[string]bool)

func CrawlConcurrent(url string, depth int, c chan []string) {
	if (regex_dont_visit.MatchString(url)) {
		fmt.Println("NOT Visiting:", url)
		c <- make([]string, 0)
		return
	}

	if _, _visited := visited[url]; _visited {
		fmt.Println("Already visited:", url)
		c <- make([]string, 0)
		return
	}

	visited[url] = true

	fmt.Println("Visiting:", url)

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("FAILED Visiting:", url, "error: ", err)
		c <- make([]string, 0)
		return
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("FAILED Read body:", url, "error: ", err)
		c <- make([]string, 0)
		return
	}

	foundUrls := FindUrls(string(body))

	for i, foundUrl := range foundUrls  {
		if (!regex_uri.MatchString(foundUrl)) {
			foundUrls[i] = url + "/" + foundUrl
		}
	}

	if depth--; depth > 0 {
		for _, foundUrl := range foundUrls  {
			c := make(chan []string)

			go CrawlConcurrent(foundUrl, depth, c)

			moreUrls := <- c

			for _, _url := range moreUrls  {
				foundUrls = append(foundUrls, _url)
			}
		}
	}

	c <- foundUrls

	return
}

func arrayStringUnique(arr []string) (arrUnique []string) {
	items := make(map[string]bool)

	for _, item := range arr  {
		items[item] = true
	}

	for item := range items {
		arrUnique = append(arrUnique, item)
	}

	return
}

func main() {
	var url string
	var depth int

	fmt.Print("URL a ser Crawleada: ")
	fmt.Scanln(&url)

	fmt.Print("Profundidade: ")
	fmt.Scanln(&depth)

	c := make(chan []string, 1)

	go CrawlConcurrent(url, depth, c)

	fmt.Println("Código concorrente rodando...")

	urls := <- c
	urls = arrayStringUnique(urls)

	for _, url := range urls  {
		fmt.Println(url)
	}

	fmt.Println(len(urls), "links found.")
}