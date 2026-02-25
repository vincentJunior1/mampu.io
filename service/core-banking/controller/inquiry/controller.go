package inquiry

import (
	utilRes "core-system/core/utils/response"
	usecaseEntities "core-system/service/core-banking/usecase/entities"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InitInquiryController(cfg ControllerConfig) (ControllerInterface, error) {
	if cfg.Usecase == nil {
		return nil, errors.New("Usecase cannot be nil")
	}

	if cfg.Log == nil {
		return nil, errors.New("Logger cannot be nil")
	}

	return &controller{
		usecase: cfg.Usecase,
		log:     cfg.Log,
	}, nil
}

func (c *controller) GetBalanceByUserID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userID"))

	if err != nil {
		c.log.Errorln("[*] Error Parse User ID: ", err)
		ctx.JSON(http.StatusBadRequest, utilRes.NewResponseWithError(http.StatusBadRequest, "Bad Request"))
		return
	}

	res := c.usecase.GetWalletByUserID(ctx, userID)
	ctx.JSON(res.StatusCode, res)
}

func (c *controller) WithdrawByUserID(ctx *gin.Context) {
	payload := usecaseEntities.WithdrawRequest{}

	if err := ctx.BindJSON(&payload); err != nil {
		c.log.Errorln("[*] Error Parse User ID: ", err)
		ctx.JSON(http.StatusBadRequest, utilRes.NewResponseWithError(http.StatusBadRequest, "Bad Request"))
		return
	}

	res := c.usecase.WithdrawByUserID(ctx, payload)

	ctx.JSON(res.StatusCode, res)
}
