package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type AddressHandler struct {
	Port int
}

func (addr *AddressHandler) index(rw http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("public/index.html")

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "Internal server error %d", http.StatusInternalServerError)
		return
	}

	tpl.Execute(rw, addr)
}

func main() {
	portPtr := flag.Int("port", 8080, "The port number")
	flag.Parse()

	router := http.NewServeMux()
	addrHandler := &AddressHandler{Port: *portPtr}

	router.HandleFunc("/", addrHandler.index)
	router.Handle("/stream", websocket.Handler(SocketServer))

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s%d", ":", *portPtr), router))
}
