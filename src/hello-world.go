package main

import (
	"fmt"
	"log"
	"net/http"
)
import "./boundary"

//!+main
func main() {
	fmt.Printf("starting hello-world server\n")
	http.HandleFunc(boundary.Handler())
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
//!-main