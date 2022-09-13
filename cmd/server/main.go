package main

import (
	"fmt"
	"log"

	"github.com/delveper/heroes/cfg"
	"github.com/delveper/heroes/core/ent"
	"github.com/delveper/heroes/core/repo"
	"github.com/delveper/heroes/core/trans"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	opt, err := cfg.NewOptions()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	kpr, err := repo.NewKeeper(opt)
	if err != nil {
		return fmt.Errorf("failed to set up repo: %w", err)
	}

	srv := ent.NewService(kpr)

	// it feels like something is missing here, and we need to redesign all
	if err := srv.CreateTable(); err != nil {
		return fmt.Errorf("failed make changes to database: %w", err)
	}

	hdl, err := trans.NewHandler(srv, opt)

	if err := hdl.Serve(); err != nil {
		return fmt.Errorf("failed to set up handler: %w", err)
	}

	return nil
}
