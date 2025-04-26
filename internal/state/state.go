package state

import "Gator/internal/config"

type State struct {
	Config *config.Config
}

// Make the state
func New(c *config.Config) *State {
	return &State{
		Config: c,
	}
}
