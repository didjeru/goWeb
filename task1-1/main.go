package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func searchTextInBodyHTML(searchText string, urls ...string) {
	for _, url := range urls {
		// Make HTTP request
		response, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		// Read response data in to memory
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal("Error reading HTTP body. ", err)
		}

		// Create a regular expression to find comments
		re := regexp.MustCompile(searchText)
		searchString := re.FindAllString(string(body), -1)
		if searchString == nil {
			fmt.Printf("No matches %s in %s\n", searchText, url)
		} else {
			for _, searchString := range searchString {
				fmt.Printf("%s - %s\n", searchString, url)
			}
		}
	}
}

func main() {
	searchTextInBodyHTML("Россия", "http://yandex.ru", "http://rambler.ru", "http://ria.ru")
	searchTextInBodyHTML("МИД", "http://yandex.ru", "http://rambler.ru", "http://ria.ru")
	searchTextInBodyHTML("Ананас", "http://yandex.ru", "http://rambler.ru", "http://ria.ru")
}
