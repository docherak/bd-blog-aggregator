package main

import (
	"log"
	"os"

	"database/sql"

	"github.com/docherak/bd-blog-aggregator/internal/config"
	"github.com/docherak/bd-blog-aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("error opening database connection: %v", err)
	}

	dbQueries := database.New(db)

	progState := state{
		config: &cfg,
		db:     dbQueries,
	}

	cmds := commands{
		commands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAggregate)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerListFollows))

	args := os.Args
	if len(args) < 2 {
		log.Fatalf("not enough arguments")
	}
	cmd := command{
		Name: args[1],
		Args: args[2:],
	}

	err = cmds.run(&progState, cmd)
	if err != nil {
		log.Fatalf("error running command '%s': %v", cmd.Name, err)
	}
}
