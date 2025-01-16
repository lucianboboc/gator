package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("login: username argument is required")
	}
	username := cmd.args[0]

	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("%s has been set\n", username)
	return nil
}
