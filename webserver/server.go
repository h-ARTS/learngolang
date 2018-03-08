package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type System struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Version      string `json:"version"`
	Serial       string `json:"serial"`
	Product      string `json:"product"`
	Model        string `json:"model"`
	GuestEnabled bool   `json:"guestEnabled"`
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		/* http.Error(w, "not found", http.StatusNotFound) */
		if r.Method == "GET" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
		w.Header().Set("Content-Type", "application/json")

		url := "http://192.168.198.48:81/api/v1/getSystemDetails"

		httpClient := http.Client{
			Timeout: time.Second * 2,
		}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}

		res, getErr := httpClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}

		data, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()

		system := System{}
		jsonErr := json.Unmarshal(data, &system)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		fmt.Fprintf(w, "%s", data)

	})

	fmt.Println("Started server on port 2000")
	log.Fatal(http.ListenAndServe(":2000", nil))
}
