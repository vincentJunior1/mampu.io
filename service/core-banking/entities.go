package corebanking

import (
	"context"
	"core-system/core/entities/server"
	"core-system/service/core-banking/controller/inquiry"

	"github.com/gin-gonic/gin"
)

type coreBankingService struct {
	engine  *gin.Engine
	inquiry inquiry.ControllerInterface
	server  server.ServerInterface
}

type CoreBankingInterface interface {
	InitRoutes() error
	Stop(ctx context.Context) error
}
