package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/arishimam/blog-aggregator/internal/config"
	"github.com/arishimam/blog-aggregator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	Db  *database.Queries
	Cfg *config.Config
}

func main() {
	args := os.Args
	//fmt.Println("ARGS: ", args)

	if len(args) < 2 {
		fmt.Println("Not enough arguments were provided")
		os.Exit(1)
	}
	if len(args) < 3 {
		fmt.Println("Username required")
		os.Exit(1)
	}

	cmd := command{args[1], args[2:]}

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	appState := state{Cfg: &cfg}
	// err = cfg.SetUser("arish")
	// if err != nil {
	// 	log.Fatalf("Couldn't set current user: %v", err)
	// }

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.run(&appState, cmd)

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	fmt.Printf("Read config again: %+v\n", cfg)

	db, err := sql.Open("postgres", cfg.DbURL)
	dbQueries := database.New(db)

	appState.Db = dbQueries

}
