// Test with curl -d 'data' localhost:9190
package main

import (
	"basic-api/handlers" // Provides the handlers package
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	l := log.New(os.Stdout, "restful-api: ", log.LstdFlags)
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

	// Go routine to start http server
	go func() {
		fmt.Printf("Starting http server on port%v\n", s.Addr)
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// Channel for go routine for handling process termination
	sigChan := make(chan os.Signal, 1)
	// signal.Notify(sigChan, os.Interrupt) -- is this an older method?
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
