package main
 
import (
    "fmt"
    "log"
    "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "OK")
}

func main() {
    http.HandleFunc("/", index)
	
	port := 10000
 
    err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
	log.Printf("start http health on port: %d", port)
}