package state

import (
	"Gator/internal/config"
	"Gator/internal/database"
)

type State struct {
	Config *config.Config
	DB     *database.Queries
}

// Make the state
func New(c *config.Config) *State {
	return &State{
		Config: c,
	}
}
