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
	router.HandleFunc("/write_cook", writeCook)
	router.HandleFunc("/read_cook", readCook)

	port := "8080"
	log.Printf("start listen on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func firstHandle(wr http.ResponseWriter, _ *http.Request) {
	_, err := wr.Write([]byte(`Hello, World!`))
	if err != nil {
		log.Print(err)
	}
}

func helloUserHandler(wr http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(wr, "Hello, %s", req.URL.Query().Get("name"))
	if err != nil {
		log.Print(err)
	}
}

func searchHandler(wr http.ResponseWriter, req *http.Request) {
	searchString := req.URL.Query().Get("string")
	_, err := fmt.Fprintf(wr, "Search string %s, found on this sites: %v", searchString, search.TextInBodyHTML(searchString))
	if err != nil {
		log.Print(err)
	}
}

func writeCook(wr http.ResponseWriter, req *http.Request) {
	http.SetCookie(wr, &http.Cookie{
		Name:    req.URL.Query().Get("name"),
		Value:   req.URL.Query().Get("value"),
		Expires: time.Now().Add(time.Minute * 10),
	})
	fmt.Fprint(wr, "Cookie changed!")
}

func readCook(wr http.ResponseWriter, req *http.Request) {
	name, err := req.Cookie(string(req.URL.Query().Get("string")))
	if err != nil {
		fmt.Fprintf(wr, "Error - ", err)
	}
	fmt.Fprintf(wr, "Cookies value - %s", name)
}
