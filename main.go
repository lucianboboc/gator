package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/lucianboboc/gator/internal/config"
	"github.com/lucianboboc/gator/internal/database"
	"os"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cmds := NewCommands()
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)

	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := state{
		db:  database.New(db),
		cfg: cfg,
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
		os.Exit(1)
	}
	commandName, args := os.Args[1], os.Args[2:]

	cmd := command{commandName, args}
	err = cmds.run(&s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
