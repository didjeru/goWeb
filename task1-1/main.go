package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func getReq(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func searchTextInBodyHTML(searchText string, urls []string) []string {
	result := make([]string, 0, 1)
	errors := 0

	for _, url := range urls {
		response, err := getReq(url)
		if err != nil {
			errors++
			log.Print(err)
			continue

		}

		if strings.Contains(string(response), searchText) {
			result = append(result, url)
		}
	}

	return result
}

func main() {

	urls := []string{
		"http://yandex.ru",
		"http://rambler.ru",
		"http://ria.ru",
	}
	log.Println(searchTextInBodyHTML("Польша", urls))
	log.Println(searchTextInBodyHTML("Россия", urls))
	log.Println(searchTextInBodyHTML("США", urls))
}
