package main

import (
	"fmt"
	"github.com/lucianboboc/gator/internal/config"
	"os"
)

type state struct {
	cfg *config.Config
}

func main() {
	cmds := NewCommands()
	cmds.register("login", handlerLogin)

	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := state{cfg: cfg}

	if len(os.Args) < 2 {
		fmt.Println("usage: command is required")
		os.Exit(1)
	}
	commandName, args := os.Args[1], os.Args[2:]
	if len(args) == 0 {
		fmt.Println("usage: login <username>")
		os.Exit(1)
	}

	cmd := command{commandName, args}
	err = cmds.run(&s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
