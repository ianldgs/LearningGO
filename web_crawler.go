package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"

	"./arrays"
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

var (
	visited   = make(map[string]bool)
	visitedMu sync.Mutex
)

func Crawl(url string, depth int, c chan []string) {
	if regex_dont_visit.MatchString(url) {
		fmt.Println("NOT Visiting:", url)
		c <- make([]string, 0)
		return
	}

	visitedMu.Lock()

	if _, _visited := visited[url]; _visited {
		fmt.Println("Already visited:", url)
		c <- make([]string, 0)
		visitedMu.Unlock()
		return
	}

	visited[url] = true

	visitedMu.Unlock()

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

	channels := make([]chan []string, len(foundUrls))

	for i, foundUrl := range foundUrls {
		channels[i] = make(chan []string, 1)

		if !regex_uri.MatchString(foundUrl) {
			//TODO: pegar o http(s)://domain da url, sem o caminho atual
			foundUrls[i] = url + "/" + foundUrl
		}
	}

	if depth--; depth >= 0 {
		for i, foundUrl := range foundUrls {
			go Crawl(foundUrl, depth, channels[i])
		}

		for i := range foundUrls {
			moreUrls := <-channels[i]

			for _, _url := range moreUrls {
				foundUrls = append(foundUrls, _url)
			}
		}
	}

	c <- foundUrls

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

	go Crawl(url, depth, c)

	fmt.Println("Código concorrente rodando...")

	urls := <-c
	urls = arrays.StringUnique(urls)

	for _, url := range urls {
		fmt.Println(url)
	}

	fmt.Println(len(urls), "links found.")

	fmt.Print("Aperte enter para finalizar...")
	fmt.Scanln()
}
