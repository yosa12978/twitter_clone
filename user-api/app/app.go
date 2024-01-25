package app

import (
	"context"
	"net/http"
	"time"

	"github.com/yosa12978/twitter/user-api/config"
	"github.com/yosa12978/twitter/user-api/db"
	"github.com/yosa12978/twitter/user-api/handlers"
	"github.com/yosa12978/twitter/user-api/logging"
)

func init() {
	config.LoadConfig()
}

type App struct {
	router http.Handler
	logger logging.Logger
	cfg    config.Config
}

func New() *App {
	app := new(App)
	app.logger = logging.New("application")
	app.router = handlers.NewRouter()
	app.cfg = config.Get()
	return app
}

func (app *App) Run(ctx context.Context) error {
	db.GetDB(ctx)

	server := http.Server{
		Addr:    app.cfg.Addr,
		Handler: handlers.NewRouter(),
	}

	errch := make(chan error, 1)
	go func() {
		app.logger.Printf("Listening @ %s", app.cfg.Addr)
		if err := server.ListenAndServe(); err != nil {
			errch <- err
		}
		close(errch)
	}()

	var err error
	select {
	case err = <-errch:
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		err = server.Shutdown(timeout)
	}
	return err
}
