package controllers

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
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
	c.nick = args[1]
	c.msg(fmt.Sprintf("ALl right i will call you %s", c.nick))
}
func (s *server) Jion(c *client, args []string) {
	RoomName := args[1]
	r, ok := s.rooms[RoomName]
	if !ok {
		r = &room{
			name:    RoomName,
			members: map[net.Addr]*client{},
		}
		s.rooms[RoomName] = r
	}
	r.members[c.conn.RemoteAddr()] = c
	if c.room != nil {

	}
	s.QuiteCurrentRoom(c)
	c.room = r
	r.BroadCast(c, fmt.Sprintf("%s has join the room", c.nick))
	c.msg(fmt.Sprintf("Welcome to %s", r.name))
}
func (s *server) Msg(c *client, args []string) {
	if c.room == nil {
		c.err(errors.New("You must join the room first.."))
		return
	}
	c.room.BroadCast(c, c.nick+": "+strings.Join(args[1:len(args)], " "))
}
func (s *server) Room(c *client, args []string) {
	var allRooms []string
	for name := range s.rooms {
		allRooms = append(allRooms, name)

	}
	c.msg(fmt.Sprintf("Available rooms are :%s", strings.Join(allRooms, ", ")))
}
func (s *server) Quit(c *client, args []string) {
	log.Println("client has disconnected %s", c.conn.RemoteAddr().String())
	s.QuiteCurrentRoom(c)
	c.msg("sad to see you go :(")
	c.conn.Close()
}
func (s *server) QuiteCurrentRoom(c *client) {
	if s.rooms != nil {
		delete(c.room.members, c.conn.RemoteAddr())
		c.room.BroadCast(c, fmt.Sprintf("%s has left from the room", c.nick))
	}
}
