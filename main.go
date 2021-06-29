package main

import (
	"github.com/CMedrado/DesafioStone/domain/account"
	"github.com/CMedrado/DesafioStone/domain/login"
	"github.com/CMedrado/DesafioStone/domain/transfer"
	https "github.com/CMedrado/DesafioStone/https"
	store_account "github.com/CMedrado/DesafioStone/store/account"
	store_login "github.com/CMedrado/DesafioStone/store/login"
	store_transfer "github.com/CMedrado/DesafioStone/store/transfer"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	time "time"
)

const dbFileName = "accounts.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal("problem opening %s %v", dbFileName, err)
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: time.RFC3339})
	lentry := logrus.NewEntry(logger)

	accountTransfer := store_transfer.NewStoredTransferAccountID()
	accountToken := store_login.NewStoredToked()
	accountStorage := store_account.NewStoredAccount(db)
	accountUseCase := account.UseCase{StoredAccount: accountStorage}
	loginUseCase := login.UseCase{AccountUseCase: &accountUseCase, StoredToken: accountToken}
	transferUseCase := transfer.UseCase{AccountUseCase: &accountUseCase, StoredTransfer: accountTransfer, TokenUseCase: &loginUseCase}
	server := https.NewServerAccount(&accountUseCase, &loginUseCase, &transferUseCase, lentry)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatal("could not hear on port 5000 ")
	}
}
