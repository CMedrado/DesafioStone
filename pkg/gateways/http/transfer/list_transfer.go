package transfer

import (
	"encoding/json"
	domain2 "github.com/CMedrado/DesafioStone/pkg/domain"
	"github.com/CMedrado/DesafioStone/pkg/domain/authentication"
	http2 "github.com/CMedrado/DesafioStone/pkg/gateways/http"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (s *Handler) ListTransfers(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	l := s.logger.WithFields(log.Fields{
		"module": "https",
		"method": "processLogin",
	})
	w.Header().Set("content-type", "application/json")
	e := errorStruct{l: l, token: token, w: w}
	accountOriginID, tokenID, err := authentication.DecoderToken(token)
	if err != nil {
		e.errorList(err)
		return
	}
	accountOrigin, err := s.account.SearchAccount(accountOriginID)
	if err != nil {
		e.errorList(err)
		return
	}
	accountToken, err := s.login.GetTokenID(tokenID)
	if err != nil {
		e.errorList(err)
		return
	}
	Transfers, err := s.transfer.GetTransfers(accountOrigin, accountToken, token)
	if err != nil {
		e.errorList(err)
		return
	}
	l.WithFields(log.Fields{
		"type": http.StatusOK,
		"time": domain2.CreatedAt(),
	}).Info("transfers handled successfully!")
	response := GetTransfersResponse{Transfers: Transfers}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (e errorStruct) errorList(err error) {
	if err != nil {
		ErrJson := http2.ErrorsResponse{Errors: err.Error()}
		if err.Error() == domain2.ErrInvalidToken.Error() {
			e.l.WithFields(log.Fields{
				"type":          http.StatusUnauthorized,
				"time":          domain2.CreatedAt(),
				"request_token": e.token,
			}).Error(err)
			e.w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(e.w).Encode(ErrJson)
		} else if err.Error() == domain2.ErrSelect.Error() {
			e.l.WithFields(log.Fields{
				"type": http.StatusInternalServerError,
				"time": domain2.CreatedAt(),
			}).Error(err)
			e.w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(e.w).Encode(ErrJson)
		} else if err.Error() == domain2.ErrInvalidID.Error() {
			e.l.WithFields(log.Fields{
				"type": http.StatusNotFound,
				"time": domain2.CreatedAt(),
			}).Error(err)
			e.w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(e.w).Encode(ErrJson)
		} else if err.Error() == domain2.ErrParse.Error() {
			e.l.WithFields(log.Fields{
				"type": http.StatusBadRequest,
				"time": domain2.CreatedAt(),
			}).Error(err)
			e.w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(e.w).Encode(ErrJson)
		} else {
			e.w.WriteHeader(http.StatusBadRequest)
		}
		return
	}
}
