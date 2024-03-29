package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "Hello, World\n")
    })
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalln(err)
    }
}
