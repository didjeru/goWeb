package search

import (
	"io/ioutil"
	"log"
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

// TextInBodyHTML - search text in a body of the HTML page
func TextInBodyHTML(searchText string, urls []string) []string {
	result := make([]string, 0, 1)

	for _, url := range urls {
		response, err := getReq(url)
		if err != nil {
			log.Print(err)
			continue
		}

		if strings.Contains(string(response), searchText) {
			result = append(result, url)
		}
	}

	return result
}
