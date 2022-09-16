package main

import (
	"fmt"
	"log"

	"github.com/delveper/heroes/cfg"
	"github.com/delveper/heroes/core/ent"
	"github.com/delveper/heroes/core/repo"
	"github.com/delveper/heroes/core/sply"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	var (
		opt *cfg.Options
		kpr *repo.Keeper
		agt *ent.Agent
		mvr *sply.Mover
		err error
	)

	if opt, err = cfg.NewOptions(); err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if kpr, err = repo.NewKeeper(opt); err != nil {
		return fmt.Errorf("failed to set up repo: %w", err)
	}

	agt = ent.NewAgent(kpr)

	// TODO: Take care of all migrations
	if err = kpr.MakeMigrations(); err != nil {
		return fmt.Errorf("failed make changes to repo: %w", err)
	}

	mvr = sply.NewMover(agt, opt)

	if err = mvr.Serve(); err != nil {
		return fmt.Errorf("failed to run logistics: %w", err)
	}

	return nil
}
