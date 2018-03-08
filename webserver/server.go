package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"Age"`
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		/* http.Error(w, "not found", http.StatusNotFound) */
		if r.Method == "GET" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
		w.Header().Set("Content-Type", "application/json")

		jsonFile, err := os.Open("data.json")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("data.json seuccessfully opened!")
		defer jsonFile.Close()

		data, _ := ioutil.ReadAll(jsonFile)

		fmt.Fprintf(w, "%s", data)
	})

	fmt.Println("Started server on port 2000")
	log.Fatal(http.ListenAndServe(":2000", nil))
}
