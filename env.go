package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("%+v\n", req)
	fmt.Fprintln(w, strings.Join(os.Environ(), "\n"))
}

func crashHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Crashing...")
	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}
	os.Exit(1)
}

// headerHandler prints out the active headers in the request
func headersHandler(w http.ResponseWriter, req *http.Request) {
	req.Header.Write(w)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/crash", crashHandler)
	http.HandleFunc("/headers", headersHandler)
	addr := ":" + os.Getenv("PORT")
	fmt.Printf("Listening on %v\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
