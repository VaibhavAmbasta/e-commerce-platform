// user-service/main.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "User Service")
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
