package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	currentUser := s.cfg.CurrentUserName
	for _, user := range users {
		name := user.Name
		if name == currentUser {
			name = fmt.Sprintf("%s (current)", user.Name)
		}
		fmt.Printf("* %s\n", name)
	}
	return nil
}
