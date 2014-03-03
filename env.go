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
	if req.URL.Path == "/crash" {
		os.Exit(1)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
