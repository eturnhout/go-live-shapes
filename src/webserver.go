package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func index(rw http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("public/index.html")

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "Internal server error %d", http.StatusInternalServerError)
		return
	}

	tpl.Execute(rw, nil)
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", index)
	router.Handle("/stream", websocket.Handler(SocketServer))

	log.Fatal(http.ListenAndServe(":8080", router))
}
