package domain

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrInvalidSecret        = errors.New("given secret is invalid")
	ErrInvalidCPF           = errors.New("given cpf is invalid")
	ErrWithoutBalance       = errors.New("given account without balance")
	ErrInvalidToken         = errors.New("given token is invalid")
	ErrInvalidID            = errors.New("given id is invalid")
	ErrInvalidAmount        = errors.New("given amount is invalid")
	ErrInvalidDestinationID = errors.New("given account destination id is invalid")
	ErrSameAccount          = errors.New("given account is the same as the account destination")
	ErrBalanceAbsent        = errors.New("given the balance amount is invalid")
	ErrLogin                = errors.New("given secret or CPF are incorrect")
	ErrAccountExists        = errors.New("given cpf is already used")
)

// CheckCPF checks if the cpf exists and returns nil if not, it returns an error
func CheckCPF(cpf string) error {

	if len(cpf) != 11 && len(cpf) != 14 {
		return errInvalidCPF
	}

	if len(cpf) == 14 {
		if string([]rune(cpf)[3]) == "." && string([]rune(cpf)[7]) == "." && string([]rune(cpf)[11]) == "-" {
			return nil
		} else {
			return errInvalidCPF
		}
	}
	return nil
}

// CheckAccountBalance checks if the account has a balance and returns nil if not, it returns an error
func CheckAccountBalance(person1 int, amount int) error {
	if person1 < amount {
		return errWithoutBalance
	}
	return nil
}

// CheckLogin Checks if the cpf and secret ar correct and returns nil if not, it returns an error
func CheckLogin(accountOrigin Account, newLogin Login) error {
	if accountOrigin.CPF != newLogin.CPF {
		return errInvalidCPF
	}
	if accountOrigin.Secret != newLogin.Secret {
		return errInvalidSecret
	}
	return nil
}

// CheckToken checks if the token is correct and returns nil if not, it returns an error
func CheckToken(token string, tokens Token) error {
	if token != tokens.Token {
		return errInvalidToken
	}
	return nil
}

// CheckExistID checks if the id exists and returns nil if not, it returns an error
func CheckExistID(accountOrigin Account) error {
	if (accountOrigin == Account{}) {
		return errInvalidID
	}
	return nil
}

// CheckAmount checks if the amount is valid and returns nil if not, it returns an error
func CheckAmount(amount int) error {
	if amount <= 0 {
		return errInvalidAmount
	}
	return nil
}

// CheckCompareID Compare two IDs to see if they are the same and returns nil if not, it returns an error
func CheckCompareID(accountOriginID, accountDestinationID uuid.UUID) error {
	if accountOriginID == accountDestinationID {
		return errSameAccount
	}
	return nil
}

// CheckExistDestinationID checks if the destination id exists and returns nil if not, it returns an error
func CheckExistDestinationID(accountOrigin Account) error {
	if (accountOrigin == Account{}) {
		return errInvalidDestinationID
	}
	return nil
}

// CheckBalance checks if the balance exists and returns nil if not, it returns an error
func CheckBalance(balance int) error {
	if balance <= 0 {
		return errBalanceAbsent
	}
	return nil
}

func CheckAccountExistence(account Account) error {
	if (account != Account{}) {
		return errAccountExists
	}
	return nil
}
