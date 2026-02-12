package inquiry

import (
	utilRes "core-system/core/utils/response"
	usecaseEntities "core-system/service/core-banking/usecase/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

func InitInquiryController(cfg ControllerConfig) (ControllerInterface, error) {
	if cfg.Usecase == nil {
		logs.Errorln("[*] Usecase cannot be nil")
	}
	return &controller{
		usecase: cfg.Usecase,
	}, nil
}

func (c *controller) GetBalanceByUserID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userID"))

	if err != nil {
		logs.Errorln("[*] Error Parse User ID: ", err)
		ctx.JSON(http.StatusBadRequest, utilRes.NewResponseWithError(http.StatusBadRequest, "Bad Request"))
		return
	}

	res := c.usecase.GetWalletByUserID(ctx, userID)
	ctx.JSON(res.StatusCode, res)
}

func (c *controller) WithdrawByUserID(ctx *gin.Context) {
	payload := usecaseEntities.WithdrawRequest{}

	if err := ctx.BindJSON(&payload); err != nil {
		logs.Errorln("[*] Error bind json: ", err)
		ctx.JSON(http.StatusBadRequest, utilRes.NewResponseWithError(http.StatusBadRequest, "Bad Request"))
		return
	}

	res := c.usecase.WithdrawByUserID(ctx, payload)

	ctx.JSON(res.StatusCode, res)
}
