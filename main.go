package main

import (
	"fmt"
	"net/http"
	"os"
)


func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "Hey Signavio, Welcome to my website!  ")
		fmt.Fprintf(w, hostname)
	})

	http.HandleFunc("/healthz", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	http.HandleFunc("/readyz", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	http.ListenAndServe(":80", nil)
}
