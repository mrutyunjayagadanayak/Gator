package command

import (
	"Gator/internal/state"
	"fmt"
)

type Commands struct {
	Handlers map[string]func(*state.State, Command) error
}

func (c *Commands) Run(s *state.State, cmd Command) error {
	handler, exist := c.Handlers[cmd.Name]
	if !exist {
		return fmt.Errorf("Command %s does not exist", cmd.Name)
	}
	return handler(s, cmd)
}

func (c *Commands) Register(name string, f func(*state.State, Command) error) {
	if c.Handlers == nil {
		c.Handlers = make(map[string]func(*state.State, Command) error)
	}
	c.Handlers[name] = f
}
