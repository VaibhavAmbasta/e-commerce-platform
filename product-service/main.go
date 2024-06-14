package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Product Service")
	})
	log.Fatal(http.ListenAndServe(":8082", nil))
}
