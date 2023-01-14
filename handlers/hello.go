package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// converts whatever you are pssing in into a http handler an register it with the server
// handlers make code Testable.
// Here we will implement dependency injection. Used for unit testing

type Hello struct {
	// handler from http pkg is a struct. It has 1 function serveHttp
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// this is a interface
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World") 		// We use log.logger instead of log. That helps in testability
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Ooops", http.StatusBadGateway)
		return
	}
	fmt.Fprintf(rw, "Hello %s", d)

}
