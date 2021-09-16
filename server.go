package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		fmt.Fprintf(w, time.Now().Format("15:04"))

	case http.MethodPost:
		fmt.Printf("Wrong method")
	}
}

func add(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	if req.Method == "POST" {
		author := req.Form.Get("author")
		message := req.Form.Get("message")
		if author != "" && message != "" {
			f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()

			_, err2 := f.WriteString(author + ":" + message + "\n")

			if err2 != nil {
				log.Fatal(err2)
			}
			fmt.Fprintf(w, author+" : "+message)
		} else {
			fmt.Println("Il faut des paramètres")
		}
	}
}

func listEntries() []string {
	raw, err := os.ReadFile("donnees.txt")
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(raw), "\n")
	return data
}

func entries(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		entries := listEntries()

		for _, rawEntry := range entries {
			entry := strings.Split(rawEntry, ":")

			fmt.Fprintf(w, entry[1]+"\n")
		}
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/add", add)
	http.HandleFunc("/entries", entries)

	http.ListenAndServe(":5768", nil)
}
