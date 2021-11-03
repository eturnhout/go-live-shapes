package main

import (
	"log"

	"golang.org/x/net/websocket"
)

var (
	// ActiveClients A map that holds the active clients.
	ActiveClients = make(map[ClientConn]int)
	// ShapesData A map that keeps track of the shapes possitions.
	ShapesData = make(map[string]ShapeData)
)

// ClientConn A client connection.
type ClientConn struct {
	websocket *websocket.Conn
	clientIP  string
}

// ShapeData the json data
type ShapeData struct {
	ID      string `json:"id"`
	YOffset int    `json:"y"`
	XOffset int    `json:"x"`
}

// SocketServer A socket server that sends and keeps track of the shape positions.
func SocketServer(ws *websocket.Conn) {
	var (
		err    error
		spData ShapeData
	)

	log.Println("Connection made by ", ws.Request().RemoteAddr)

	if len(ShapesData) > 0 {
		for shapeID := range ShapesData {
			websocket.JSON.Send(ws, ShapesData[shapeID])
		}
	}

	cliAddr := ws.Request().RemoteAddr
	cliConn := ClientConn{
		websocket: ws,
		clientIP:  cliAddr,
	}

	ActiveClients[cliConn] = 0

	for {
		if err = websocket.JSON.Receive(ws, &spData); err != nil {
			log.Println("Connection closed by ", ws.Request().RemoteAddr)
			delete(ActiveClients, cliConn)
			return
		}

		ShapesData[spData.ID] = spData

		for cn := range ActiveClients {
			go func(cn ClientConn) {
				if cn.clientIP != ws.Request().RemoteAddr {
					if err = websocket.JSON.Send(cn.websocket, spData); err != nil {
						log.Println("Unable to send data to ", cn.clientIP)
					}
				}
			}(cn)
		}
	}
}
