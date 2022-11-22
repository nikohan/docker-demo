package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func get(w http.ResponseWriter, r *http.Request) {
	res, _ := http.Get("https://www.google.com/")
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Fprintf(w, string(body))
}

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/get", get)

	port := 80

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Printf("start http health on port: %d", port)
}
