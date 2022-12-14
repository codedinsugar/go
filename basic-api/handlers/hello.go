package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	h.l.Println("Hello from handlers package") // Consuming Hello struct and NewHello method

	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Uh oh something didn't go right...", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(res, "Hello %s\n", data)
	log.Printf("Data: %s\n", data)
}
