package controllers

import "net"

type client struct {
	conn     net.Conn
	nick     string
	room     *room
	commands chan<- commandID
}
