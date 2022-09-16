package sply

import (
	"fmt"
	"net/http"

	"github.com/delveper/heroes/cfg"
)

// Mover is responsible for carrying logistics
type Mover struct {
	Agent   UserMover
	Options *cfg.Options
	Router  *http.ServeMux
	Server  *http.Server
}

// NewMover returns a pointer receiver to Mover
func NewMover(um UserMover, opt *cfg.Options) *Mover {
	mvr := &Mover{
		Agent:   um,
		Options: opt,
	}

	mvr.Router = http.NewServeMux()

	mvr.Server = &http.Server{
		Handler:      mvr.Router,
		Addr:         mvr.Options.HTTP.Host + ":" + mvr.Options.HTTP.Port,
		ReadTimeout:  mvr.Options.HTTP.ReadTimeout,
		WriteTimeout: mvr.Options.HTTP.WriteTimeout,
		IdleTimeout:  mvr.Options.HTTP.IdleTimeout,
	}

	mvr.RegisterRoutes()

	return mvr
}

func (mvr *Mover) RegisterRoutes() {
	mvr.Router.Handle(mvr.Options.API.User.Endpoint, // pattern
		Wrap(mvr.HandleUser(), // main handler for the endpoint that will be http method aware
			WithJSON, // here we can insert e all king of middleware
		))
	// to be continue...
}

func (mvr *Mover) Serve() error {
	if err := mvr.Server.ListenAndServe(); err != nil {
		return fmt.Errorf("error loading the server: %w", err)
	}
	return nil
}
