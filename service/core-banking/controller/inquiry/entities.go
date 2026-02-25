package inquiry

import (
	usecase "core-system/service/core-banking/usecase/inquiry"

	"core-system/core/entities/logger"

	"github.com/gin-gonic/gin"
)

type ControllerInterface interface {
	GetBalanceByUserID(ctx *gin.Context)
	WithdrawByUserID(ctx *gin.Context)
}

type controller struct {
	usecase usecase.UsecaseInterface
	log     logger.LoggerInterface
}

type ControllerConfig struct {
	Usecase usecase.UsecaseInterface
	Log     logger.LoggerInterface
}
