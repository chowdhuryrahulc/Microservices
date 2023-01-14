package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chowdhuryrahulc/microservices/handlers"
)

// Structure from part2. to make it Testable
// Step1: put all the routing code into a struct
func main()  {
	// we will write our logs in os.Stdout
	// "product-api" is our prefix in all our logs
	// func log.New(out io.Writer, prefix string, flag int) *log.Logger
	//todo What are flags?
	l := log.New(os.Stdout, "product-api", log.LstdFlags)		
	hh := handlers.NewHello(l)

	// we convert function to a handler type, and then registering it into defaultserve mux
	sm := http.NewServeMux()
	sm.Handle("/", hh)

	//! Left bcoz too complicated
}
