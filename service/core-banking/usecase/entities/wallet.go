package entities

type WalletResponse struct {
	ID      int     `json:"id"`
	Balance float64 `json:"balance"`
}

type WithdrawRequest struct {
	UserID int     `json:"userID"`
	Amount float64 `json:"amount"`
}
