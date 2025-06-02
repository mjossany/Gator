package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mjossany/Gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: &s <name>", cmd.Name)
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	})

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	s.cfg.SetUser(user.Name)
	fmt.Println("User successfully created", user)

	return nil
}
