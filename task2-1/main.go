package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"./search"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", firstHandle)
	router.HandleFunc("/user", helloUserHandler)
	router.HandleFunc("/search", searchHandler)
	router.HandleFunc("/write_cookie", writeCookieHandler)
	router.HandleFunc("/read_cookie", readCookieHandler)

	port := "8080"
	log.Printf("start listen on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func firstHandle(wr http.ResponseWriter, _ *http.Request) {
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
	searchString := req.URL.Query().Get("string")
	if _, err := fmt.Fprintf(wr, "Search string %s, found on this sites: %v", searchString, search.TextInBodyHTML(searchString)); err != nil {
		log.Print(err)
	}
}

func writeCookieHandler(wr http.ResponseWriter, req *http.Request) {
	http.SetCookie(wr, &http.Cookie{
		Name:    req.URL.Query().Get("name"),
		Value:   req.URL.Query().Get("value"),
		Expires: time.Now().Add(time.Minute * 10),
	})
	fmt.Fprint(wr, "Cookie changed!")
}

func readCookieHandler(wr http.ResponseWriter, req *http.Request) {
	name, err := req.Cookie(string(req.URL.Query().Get("string")))
	if err != nil {
		fmt.Fprint(wr, "Error - ", err)
	}
	fmt.Fprintf(wr, "Cookies value - %s", name)
}
