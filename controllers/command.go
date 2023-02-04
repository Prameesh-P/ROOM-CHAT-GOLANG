package controllers

type commandID int

const (
	CMD_NICK commandID = iota
	CMD_JION
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
)

type command struct {
	id     commandID
	client *client
	args   []string
}
