package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "GETメソッド-HOTRELOAD")
		case http.MethodPost:
			fmt.Fprintf(w, "POSTメソッド")
		case http.MethodPut:
			fmt.Fprintf(w, "PUTメソッド")
		case http.MethodDelete:
			fmt.Fprintln(w, "DELETEメソッド")
		default:
			http.Error(w, "メソッドがサポートされていません", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}