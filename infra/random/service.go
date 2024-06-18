package random

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"sync"
	"x-bank-users/cerrors"
	"x-bank-users/ercodes"
)

type Service struct {
	randomBytes []byte
	mu          *sync.Mutex
}

func NewService() Service {
	return Service{
		randomBytes: make([]byte, 2),
		mu:          &sync.Mutex{},
	}
}

func (s *Service) GenerateString(_ context.Context, set string, size int) (string, error) {
	res := make([]byte, size)
	for i := 0; i < size; i++ {
		randomNum, err := s.GenerateRandomNum()
		if err != nil {
			return "", err
		}

		res[i] = set[int(randomNum)%len(set)]
	}

	return string(res), nil
}

func (s *Service) GenerateRandomNum() (uint16, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := rand.Read(s.randomBytes)
	if err != nil {
		return 0, cerrors.NewErrorWithUserMessage(ercodes.RandomGeneration, err, "Ошибка генерации случайного числа")
	}

	num := binary.BigEndian.Uint16(s.randomBytes)

	return num, nil
}
