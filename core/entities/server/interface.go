package server

import "context"

type ServerInterface interface {
	StartServer() error
	ShutdownServer(ctx context.Context) error
}
