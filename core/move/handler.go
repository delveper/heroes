package move

import (
	"fmt"
	"net/http"

	"github.com/delveper/heroes/cfg"
)

// Handlers are responsible for carrying the logic
// of writing response headers and bodies

type Lug struct {
	Router  *http.ServeMux
	Service UserService
	Server  *http.Server
	Options *cfg.Options
}

// NewLug returns a pointer to Lug
// or non nil err in worst case scenario
func NewLug(srv UserService, opt *cfg.Options) (lug *Lug, err error) {
	lug = &Lug{
		Service: srv,
		Options: opt,
	}

	lug.Router = http.NewServeMux()

	lug.Server = &http.Server{
		Handler:      lug.Router,
		Addr:         lug.Options.HTTP.Host + ":" + lug.Options.HTTP.Port,
		ReadTimeout:  lug.Options.HTTP.ReadTimeout,
		WriteTimeout: lug.Options.HTTP.WriteTimeout,
		IdleTimeout:  lug.Options.HTTP.IdleTimeout,
	}

	lug.RegisterRoutes()

	return lug, nil
}

func (lug *Lug) RegisterRoutes() {
	lug.Router.Handle(lug.Options.API.User.Endpoint, // pattern
		Wrap(lug.HandleUser(), // main handler for the endpoint that will be http method aware
			WithJSON, // here we can insert e all king of middleware
		))
}

func (lug *Lug) Serve() error {
	if err := lug.Server.ListenAndServe(); err != nil {
		return fmt.Errorf("error loading the server: %w", err)
	}
	return nil
}
