package controllers

import (
	"log"
	"net"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func NewServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) NewClient(conn net.Conn) {
	log.Println("new client is connected..%s", conn.RemoteAddr().String())
	var c = &client{
		conn:    conn,
		nick:    "anonymous",
		room:    nil,
		command: s.commands,
	}
	c.ReadInput()
}
func (s *server) Run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.Nick(cmd.client, cmd.args)
		case CMD_JION:
			s.Jion(cmd.client, cmd.args)
		case CMD_MSG:
			s.Msg(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.Room(cmd.client, cmd.args)
		case CMD_QUIT:
			s.Quit(cmd.client, cmd.args)
		}
	}
}
func (s *server) Nick(c *client, args []string) {

}
func (s *server) Jion(c *client, args []string) {

}
func (s *server) Msg(c *client, args []string) {

}
func (s *server) Room(c *client, args []string) {

}
func (s *server) Quit(c *client, args []string) {

}
