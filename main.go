package main

import (
	"fmt"
	"log"

	"github.com/docherak/bd-blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	err = cfg.SetUser("docherak")
	if err != nil {
		log.Fatalf("couldn't set current user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error loading updated config: %v", err)
	}

	fmt.Printf("%+v", cfg)

}
