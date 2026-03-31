package main

import (
	"log"
	"os"

	"github.com/vmarin93/gator/internal/config"
)

type state struct {
	conf *config.Config
}

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}
	appState := &state{
		conf: &conf,
	}
	cmds := commands{
		validCmds: map[string]func(*state, command) error{},
	}
	cmds.register("login", HandlerLogin)
	if len(os.Args) < 2 {
		log.Fatal("Usage: gator <command> [args...]")
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	if err := cmds.run(appState, command{name: cmdName, args: cmdArgs}); err != nil {
		log.Fatal(err)
	}
}
