package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})

	fs := http.FileServer(http.Dir("src/"))
	http.Handle("/src/", http.StripPrefix("/src/", fs))

	http.ListenAndServe(":2000", fs)
}
