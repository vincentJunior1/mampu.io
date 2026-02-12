package usecase

import (
	"core-system/core/entities/response"
	resUtil "core-system/core/utils/response"
	"core-system/service/core-banking/usecase/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (u *usecase) GetWalletByUserID(ctx *gin.Context, userID int) response.Response {
	wallet, err := u.repo.GetBalanceByUserID(ctx, userID)

	if err != nil {
		return resUtil.NewResponseWithError(http.StatusOK, "Wallet Not Found")
	}

	res := entities.WalletResponse{
		ID:      wallet.ID,
		Balance: wallet.Balance,
	}

	return resUtil.NewResponse(http.StatusOK, "Success", res)
}

func (u *usecase) WithdrawByUserID(ctx *gin.Context, payload entities.WithdrawRequest) response.Response {
	wallet, err := u.repo.WithdrawByUserID(ctx, payload)

	if err != nil {
		logrus.Errorln("[*] Error Withdraw: ", err)
		return resUtil.NewResponseWithError(http.StatusBadRequest, "Unprocessable Entity")
	}

	return resUtil.NewResponse(http.StatusOK, "Success", wallet)
}
