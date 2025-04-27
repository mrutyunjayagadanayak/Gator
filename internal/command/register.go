package command

import (
	"Gator/internal/database"
	"Gator/internal/state"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

func HandlerRegister(s *state.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("no username given")
	}
	user, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      cmd.Args[0],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return fmt.Errorf("user %s already exists", cmd.Args[0])
		}
		return fmt.Errorf("error creating user - %v", err)
	}

	s.Config.SetUser(user.Name)
	fmt.Printf("User %s created with ID %s\n", user.Name, user.ID)
	return nil
}
