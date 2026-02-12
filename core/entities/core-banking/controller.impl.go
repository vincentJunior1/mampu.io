package corebanking

import "github.com/gin-gonic/gin"

type CoreBankingInterface interface {
	GetByID(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Post(ctx *gin.Context)
}
