package inquiry

import (
	usecase "core-system/service/core-banking/usecase/inquiry"

	"github.com/gin-gonic/gin"
)

type ControllerInterface interface {
	GetBalanceByUserID(ctx *gin.Context)
	WithdrawByUserID(ctx *gin.Context)
}

type controller struct {
	usecase usecase.UsecaseInterface
}

type ControllerConfig struct {
	Usecase usecase.UsecaseInterface
}
