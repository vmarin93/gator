package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	validCmds map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	cmdFunc, ok := c.validCmds[cmd.name]
	if !ok {
		err := fmt.Errorf("Command %s is not a valid command", cmd.name)
		return err
	}
	return cmdFunc(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.validCmds[name] = f
}
