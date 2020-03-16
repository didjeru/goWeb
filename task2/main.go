package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"./search"
	"github.com/google/uuid"
)

// SearchValues - struct of search values
type SearchValues struct {
	String string   `json:"search"`
	Sites  []string `json:"sites"`
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", startPageHandler)
	router.HandleFunc("/user", helloUserHandler)
	router.HandleFunc("/search", searchHandler)
	router.HandleFunc("/write_cookie", writeCookieHandler)
	router.HandleFunc("/read_cookie", readCookieHandler)

	port := "8080"
	log.Printf("start listen on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func startPageHandler(wr http.ResponseWriter, _ *http.Request) {
	if _, err := wr.Write([]byte(`Hello, World!`)); err != nil {
		log.Print(err)
	}
}

func helloUserHandler(wr http.ResponseWriter, req *http.Request) {
	if _, err := fmt.Fprintf(wr, "Hello, %s", req.URL.Query().Get("name")); err != nil {
		log.Print(err)
	}
}

func searchHandler(wr http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprint(wr, "Request should be POST")
		return
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print(err)
		return
	}
	searchValues := new(SearchValues)
	if err := json.Unmarshal(body, searchValues); err != nil {
		log.Print(err)
		return
	}
	result := search.TextInBodyHTML(searchValues.String, searchValues.Sites)
	bytes, err := json.Marshal(result)
	if err != nil {
		log.Print(err)
		return
	}
	_, err = wr.Write(bytes)
	if err != nil {
		log.Print(err)
		return
	}
}

func writeCookieHandler(wr http.ResponseWriter, req *http.Request) {
	if sessionToken, err := uuid.NewUUID(); err != nil {
		fmt.Fprint(wr, err)
	} else {
		http.SetCookie(wr, &http.Cookie{
			Name:    "session_token",
			Value:   sessionToken.String(),
			Expires: time.Now().Add(time.Minute * 10),
		})
		fmt.Fprint(wr, "Cookie changed!")
	}
}

func readCookieHandler(wr http.ResponseWriter, req *http.Request) {
	name, err := req.Cookie(string(req.URL.Query().Get("string")))
	if err != nil {
		fmt.Fprint(wr, "Error - ", err)
		return
	}
	fmt.Fprintf(wr, "Cookies value - %s", name)
}
