package ercodes

import "x-bank-users/cerrors"

const (
	_ cerrors.Code = -iota

	UserNotFound
	ActivationCodeNotFound
	BcryptHashing
	RandomGeneration
	HS512Authorization
	RS256Authorization
	AccountNotActivated
	WrongPassword
	Invalid2FACode
	InvalidEmailOrLogin
)
