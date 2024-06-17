package jwt

import (
	"context"
	"x-bank-users/auth"
)

type (
	RS256 struct {
	}
)

func NewRS256() (RS256, error) {
	return RS256{}, nil
}

func (R *RS256) Authorize(ctx context.Context, claims auth.Claims) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (R *RS256) VerifyAuthorization(ctx context.Context, authorization []byte) (auth.Claims, error) {
	//TODO implement me
	panic("implement me")
}
