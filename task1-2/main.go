package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const (
	yandexLogin      = ""
	yandexPassword   = ""
	yandexPathToFile = "https://yadi.sk/i/cwPl5546eUNUuQ"
)

type yandexDiskResponse struct {
	Type string `json:type`
	File string `json:file`
	Name string `json:name`
}

func getJSON(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// return json.NewDecoder(resp.Body).Decode(target)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, target)
}

func downloadFileToHardDisk(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath, body, 0644)
	return err

}

func downloadFileFromYandexDisk(file string) (string, error) {
	result := new(yandexDiskResponse)
	publicKey := url.QueryEscape(file)
	downloadURL := "https://cloud-api.yandex.net/v1/disk/public/resources?public_key=" + publicKey

	if err := getJSON(downloadURL, result); err != nil {
		return "", err
	}

	if result.Type != "file" {
		return "", fmt.Errorf("This is not file")
	}

	if err := downloadFileToHardDisk(result.Name, result.File); err != nil {
		return "", err
	}

	return result.Name, nil
}

func uploadFileToYandexDiskFromHardDisk(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	request, err := http.NewRequest(http.MethodPut, "https://webdav.yandex.ru/"+filepath, file)
	if err != nil {
		return err
	}

	request.SetBasicAuth(yandexLogin, yandexPassword)

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return fmt.Errorf("Status is: " + resp.Status)
	}
	return nil
}

func main() {
	filePath, err := downloadFileFromYandexDisk(yandexPathToFile)
	if err != nil {
		panic(err.Error())
	}

	if err := uploadFileToYandexDiskFromHardDisk(filePath); err != nil {
		panic(err.Error())
	}

	fmt.Println("All done")
}
