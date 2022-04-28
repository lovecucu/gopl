package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fetch(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func fetch(w io.Writer) {
	if len(os.Args) <= 1 {
		fmt.Fprintf(w, "fetch: lack of target url")
		return
	}
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(w, "fetch: %v\n", err)
		return
	}
	if _, err := io.Copy(w, resp.Body); err != nil {
		fmt.Fprintf(w, "fetch: reading %s: %v\n", url, err)
		return
	}
}
