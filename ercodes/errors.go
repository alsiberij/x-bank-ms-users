package ercodes

import "x-bank-users/cerrors"

const (
	_ cerrors.Code = -iota

	UserNotFound
	ActivationCodeNotFound
	BcryptHashing
	RandomGeneration
)
