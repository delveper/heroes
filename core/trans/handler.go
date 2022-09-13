package trans

import (
	"fmt"
	"net/http"

	"github.com/delveper/heroes/cfg"
)

// Handlers are responsible for carrying the logic
// of writing response headers and bodies

type Handler struct {
	Router  *http.ServeMux
	Service UserService
	Server  *http.Server
	Options *cfg.Options
}

// NewHandler returns a pointer to Handler
// or non nil err in worst case scenario
func NewHandler(srv UserService, opt *cfg.Options) (hdl *Handler, err error) {
	hdl = &Handler{
		Service: srv,
		Options: opt,
	}

	hdl.Router = http.NewServeMux()

	hdl.Server = &http.Server{
		Handler:      hdl.Router,
		Addr:         hdl.Options.HTTP.Host + ":" + hdl.Options.HTTP.Port,
		ReadTimeout:  hdl.Options.HTTP.ReadTimeout,
		WriteTimeout: hdl.Options.HTTP.WriteTimeout,
		IdleTimeout:  hdl.Options.HTTP.IdleTimeout,
	}

	hdl.RegisterRoutes()

	return hdl, nil
}

func (hdl *Handler) RegisterRoutes() {
	hdl.Router.Handle(hdl.Options.API.User.Endpoint, // pattern
		Wrap(hdl.UserHandle(), // main handler for the endpoint that will be http method aware
			WithJSON, // here we can insert e all king of middleware
		))
}

func (hdl *Handler) Serve() error {
	if err := hdl.Server.ListenAndServe(); err != nil {
		return fmt.Errorf("error loading the server: %w", err)
	}
	return nil
}
