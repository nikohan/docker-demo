package main
 
import (
    "fmt"
    "log"
    "net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "OK")
}

func main() {
    http.HandleFunc("/health", health)
	
	port := 80
 
    err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
	log.Printf("start http health on port: %d", port)
}