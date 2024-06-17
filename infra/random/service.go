package random

import (
	"context"
	"crypto/rand"
	"sync"
	"x-bank-users/cerrors"
	"x-bank-users/ercodes"
)

type Service struct {
	buf []byte
	res []byte
	mu  *sync.Mutex
}

func NewService() Service {
	return Service{
		mu: &sync.Mutex{},
	}
}

func (s *Service) GenerateString(_ context.Context, set string, size int) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.res = nil
	s.buf = make([]byte, size)

	_, err := rand.Read(s.buf)
	if err != nil {
		return "", cerrors.NewErrorWithUserMessage(ercodes.RandomGeneration, err, "Ошибка генерации строки")
	}

	for _, b := range s.buf {
		s.res = append(s.res, set[int(b)%len(set)])
	}

	return string(s.res), nil
}
