package jwt

import (
	"context"
	"x-bank-users/auth"
)

type (
	HS512 struct {
	}
)

func NewHS512() (HS512, error) {
	return HS512{}, nil
}

func (R *HS512) Authorize(ctx context.Context, claims auth.Claims) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (R *HS512) VerifyAuthorization(ctx context.Context, authorization []byte) (auth.Claims, error) {
	//TODO implement me
	panic("implement me")
}
