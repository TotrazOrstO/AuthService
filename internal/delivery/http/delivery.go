package http

import (
	"MedodsProject/internal/user"
	"MedodsProject/pkg/config"
	"context"
	"fmt"
	"net/http"
)

type Delivery struct {
	server *http.Server
	user   *user.Service
}

func New(cfg config.HTTP, userService *user.Service) *Delivery {
	d := &Delivery{
		user: userService,
	}

	mux := d.initRoutes()
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler: mux,
	}

	d.server = server

	return d
}

func (d *Delivery) Start(ctx context.Context) error {
	return d.server.ListenAndServe()
}

func (d *Delivery) Close() error {
	return d.server.Close()
}

func (d *Delivery) initRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/access", d.access)
	mux.HandleFunc("/refresh", d.refresh)

	return mux
}
