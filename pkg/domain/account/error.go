package account

import "errors"

var (
	ErrBalanceAbsent = errors.New("given the balance amount is invalid")
	ErrUpdate        = errors.New("unable to update")
)
