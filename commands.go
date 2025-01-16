package main

import "errors"

type command struct {
	name string
	args []string
}

func NewCommands() *commands {
	return &commands{
		commandsMap: make(map[string]func(*state, command) error),
	}
}

type commands struct {
	commandsMap map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandsMap[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.commandsMap[cmd.name]
	if !ok {
		return errors.New("command not found: " + cmd.name)
	}
	return f(s, cmd)
}
