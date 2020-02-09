package boundary

import (
	"fmt"
	"log"
	"net/http"
)

const (
	helloPath = "/hello"
)

func Handler() (string, func(w http.ResponseWriter, req *http.Request)) {
	return helloPath, execute
}

func execute(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling request method %s %s", req.Method, req.URL)
	_, _ = fmt.Fprintf(w, "hello world\n")
}


