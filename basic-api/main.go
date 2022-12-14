// Test with curl -d 'data' localhost:9190
package main

import (
	"basic-api/handlers" // Provides the handlers package
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	// https://pkg.go.dev/net/http#NewServeMux
	// Our own implementation of the DefaultServeMux
	// Our serve mux will handle all handler paths
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	// https://pkg.go.dev/net/http#Server
	// Our own implementation of the default http server
	// We define the parameters for running an http server
	s := &http.Server{
		Addr:         ":9190",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	s.ListenAndServe() // Start listening for http calls
}
