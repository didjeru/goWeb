package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

type yandexDiskResponse struct {
	Type string `json:type`
	File string `json:file`
	Name string `json:name`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(url string, target interface{}) error {
	resp, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func downloadFileFromYandexDisk(file string) error {
	result := new(yandexDiskResponse)
	publicKey := url.QueryEscape(file)
	downloadURL := "https://cloud-api.yandex.net/v1/disk/public/resources?public_key=" + publicKey

	if err := getJSON(downloadURL, result); err != nil {
		return err
	}

	if result.Type != "file" {
		return fmt.Errorf("This is not file")
	}

	if err := downloadFile(result.Name, result.File); err != nil {
		return err
	}

	return nil
}

func main() {
	 //only download
	if err := downloadFileFromYandexDisk("https://yadi.sk/i/cwPl5546eUNUuQ"); err != nil {
		panic(err.Error())
	}
}
