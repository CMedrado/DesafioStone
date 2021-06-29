package https

import (
	"github.com/CMedrado/DesafioStone/domain"
	store_account "github.com/CMedrado/DesafioStone/store/account"
	store_transfer "github.com/CMedrado/DesafioStone/store/transfer"
	"github.com/sirupsen/logrus"
	"net/http"
)

type CreateRequest struct {
	Name    string `json:"name"`
	CPF     string `json:"cpf"`
	Secret  string `json:"secret"`
	Balance int    `json:"balance"`
}

type CreateResponse struct {
	ID int `json:"id"`
}

type BalanceResponse struct {
	Balance int `json:"balance"`
}

type LoginRequest struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type TransfersRequest struct {
	AccountOriginID      int `json:"account_origin_id"`
	Token                int `json:"token"`
	AccountDestinationID int `json:"account_destination_id"`
	Amount               int `json:"amount"`
}

type TransferResponse struct {
	ID int `json:"id"`
}

type GetTransfersResponse struct {
	Transfers []store_transfer.Transfer `json:"transfers"`
}

type GetAccountsResponse struct {
	Accounts []store_account.Account `json:"accounts"`
}

type ServerAccount struct {
	account  domain.AccountUseCase
	login    domain.LoginUseCase
	transfer domain.TransferUseCase
	logger   *logrus.Entry

	http.Handler
}

type ErrorsResponse struct {
	Errors string `json:"errors"`
}
