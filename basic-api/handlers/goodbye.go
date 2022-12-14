package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	g.l.Println("Hello from handlers package") // Consuming Goodbye struct and NewGoodbye method

	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Uh oh something didn't go right...", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(res, "Goodbye %s\n", data)
	log.Printf("Data: %s\n", data)
	res.Write([]byte("Adios muchachos"))

}
