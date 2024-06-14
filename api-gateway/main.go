package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func reverseProxy(target string) http.Handler {
	url, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(url)
}

func main() {
	http.Handle("/user", reverseProxy("http://localhost:8081"))
	http.Handle("/product", reverseProxy("http://localhost:8082"))
	http.Handle("/order", reverseProxy("http://localhost:8083"))
	http.Handle("/payment", reverseProxy("http://localhost:8084"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API Gateway")
	})

	log.Println("API Gateway running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
