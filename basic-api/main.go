// Test with curl -d 'data' localhost:9190
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//HandleFunc is a convenience method that creates an http handler to respond to http requests
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(res, "Uh oh something didn't go right...", http.StatusBadRequest)
		} else {
			fmt.Fprintf(res, "Hello %s\n", data)
			log.Printf("Data: %s\n", data)
		}
	})

	http.HandleFunc("/goodbye", func(res http.ResponseWriter, req *http.Request) {
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(res, "Uh oh something didn't go right...", http.StatusBadRequest)
		} else {
			fmt.Fprintf(res, "Goodbye %s\n", data)
			log.Printf("Data: %s\n", data)
		}
	})

	http.ListenAndServe(":9190", nil) // 9090 is taken by Cockpit UI
}
