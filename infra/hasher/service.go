package hasher

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"x-bank-users/cerrors"
	"x-bank-users/ercodes"
)

type (
	Service struct {
	}
)

func NewService() Service {
	return Service{}
}

func (s *Service) HashPassword(_ context.Context, password []byte, cost int) ([]byte, error) {
	passwordHash, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		return nil, cerrors.NewErrorWithUserMessage(ercodes.BcryptHashing, err, "Ошибка хэширования пароля")
	}

	return passwordHash, nil
}
