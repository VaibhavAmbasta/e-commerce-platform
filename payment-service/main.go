package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Payment Service")
	})
	log.Fatal(http.ListenAndServe(":8084", nil))
}
