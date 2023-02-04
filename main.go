package main

import (
	"github.com/Prameesh-P/ROOM-CHAT-GOLANG/controllers"
	"log"
	"net"
)

func main() {
	server := controllers.NewServer()
	go server.Run()
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("Unable to connect with the server %s", err.Error())
	}
	defer listener.Close()
	log.Fatal("started server on port :8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("unable to accept to connection %s", err.Error())
			continue
		}
		go server.NewClient(conn)
	}
}
