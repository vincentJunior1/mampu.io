package server

import (
	entitiesServer "core-system/core/entities/server"
	"net/http"
	"time"
)

func InitServer(cfg entitiesServer.ServerConfig) (entitiesServer.ServerInterface, error) {
	return &config{
		port:         cfg.Port,
		handler:      cfg.Handler,
		idleTimeout:  cfg.IdleTimeout,
		readTimeout:  cfg.ReadTimeout,
		writeTimeout: cfg.WriteTimeout,
	}, nil
}

func (c *config) StartServer() error {
	srv := http.Server{
		Addr:           c.port,
		Handler:        c.handler,
		IdleTimeout:    time.Duration(c.idleTimeout) * time.Second,
		ReadTimeout:    time.Duration(c.readTimeout) * time.Second,
		WriteTimeout:   time.Duration(c.writeTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := srv.ListenAndServe()
	return err
}
