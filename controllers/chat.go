package controllers

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn    net.Conn
	nick    string
	room    *room
	command chan<- command
}

func (c *client) ReadInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}
		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])
		switch cmd {
		case "/nick":
			c.command <- command{
				id:     CMD_NICK,
				client: c,
				args:   args,
			}
		case "/join":
			c.command <- command{
				id:     CMD_JION,
				client: c,
				args:   args,
			}
		case "/rooms":
			c.command <- command{
				id:     CMD_ROOMS,
				client: c,
				args:   args,
			}
		case "/msg":
			c.command <- command{
				id:     CMD_MSG,
				client: c,
				args:   args,
			}
		case "/quit":
			c.command <- command{
				id:     CMD_QUIT,
				client: c,
				args:   args,
			}
		default:
			c.err(fmt.Errorf("unkwon command"))
		}
	}
}
func (c *client) err(err error) {
	c.conn.Write([]byte("ERR:" + err.Error() + "\n"))
}
func (c *client) msg(msg string) {
	c.conn.Write([]byte("> " + msg + "\n"))
}
