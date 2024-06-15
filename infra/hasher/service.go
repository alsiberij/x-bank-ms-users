package hasher

import (
	"context"
	"golang.org/x/crypto/bcrypt"
)

type (
	Service struct {
	}
)

func NewService() Service {
	return Service{}
}

func (s *Service) HashPassword(_ context.Context, password []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, cost)
}
