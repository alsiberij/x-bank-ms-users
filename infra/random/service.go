package random

import (
	"context"
	"crypto/rand"
	"math/big"
	"x-bank-users/cerrors"
	"x-bank-users/ercodes"
)

type Service struct {
}

func NewService() Service {
	return Service{}
}

func (s *Service) GenerateString(_ context.Context, set string, size int) (string, error) {
	res := make([]byte, size)
	for i := 0; i < size; i++ {
		randomNum, err := s.GenerateRandomNum(size)
		if err != nil {
			return "", cerrors.NewErrorWithUserMessage(ercodes.RandomGeneration, err, "Ошибка генерации случайной строки")
		}
		res[i] = set[randomNum]
	}
	return string(res), nil
}

func (s *Service) GenerateRandomNum(n int) (int, error) {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		return 0, cerrors.NewErrorWithUserMessage(ercodes.RandomGeneration, err, "Ошибка генерации случайного числа")
	}
	return int(num.Int64()), nil
}
