package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/vmarin93/gator/internal/config"
	"github.com/vmarin93/gator/internal/database"
)

type state struct {
	db   *database.Queries
	conf *config.Config
}

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}
	db, err := sql.Open("postgres", conf.DBURL)
	if err != nil {
		log.Fatalf("error opening db connection: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)
	appState := &state{
		db:   dbQueries,
		conf: &conf,
	}
	cmds := commands{
		validCmds: map[string]func(*state, command) error{},
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handleReset)
	cmds.register("users", handlerListUsers)
	if len(os.Args) < 2 {
		log.Fatal("Usage: gator <command> [args...]")
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	if err := cmds.run(appState, command{name: cmdName, args: cmdArgs}); err != nil {
		log.Fatal(err)
	}
}
