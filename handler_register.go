package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lucianboboc/gator/internal/database"
	"time"
)

var (
	ErrUserAlreadyExists = errors.New("User already exists")
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("invalid arguments")
	}

	now := time.Now()
	p := database.CreateUserParams{
		ID:        uuid.New(),
		Name:      cmd.args[0],
		CreatedAt: now,
		UpdatedAt: now,
	}
	user, err := s.db.CreateUser(context.Background(), p)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Println("User created:", user.Name)
	return nil
}
