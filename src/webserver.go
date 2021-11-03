package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type Address struct {
	Host string
}

func index(rw http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("public/index.html")

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "Internal server error %d", http.StatusInternalServerError)
		return
	}

	tpl.Execute(rw, Address{Host: r.Host})
}

func main() {
	portPtr := flag.Int("port", 8080, "The port number")
	router := http.NewServeMux()

	flag.Parse()

	router.HandleFunc("/", index)
	router.Handle("/stream", websocket.Handler(SocketServer))

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s%d", ":", *portPtr), router))
}
