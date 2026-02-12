package usecase

import (
	"core-system/core/entities/response"
	"core-system/service/core-banking/repository"
	"core-system/service/core-banking/usecase/entities"

	"github.com/gin-gonic/gin"
)

type UsecaseInterface interface {
	GetWalletByUserID(ctx *gin.Context, userID int) response.Response
	WithdrawByUserID(ctx *gin.Context, payload entities.WithdrawRequest) response.Response
}

type UsecaseConfig struct {
	Repo repository.RepoInterface
}

type usecase struct {
	repo repository.RepoInterface
}
