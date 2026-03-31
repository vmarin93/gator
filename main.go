package main

import (
	"fmt"
	"log"

	"github.com/vmarin93/gator/internal/config"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}
	fmt.Printf("Read config: %+v\n", conf)
	if err := conf.SetUser("vasile"); err != nil {
		log.Fatalf("couldn't set current user: %v", err)
	}
	conf, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", conf)
}
