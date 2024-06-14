package swissknife

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"sync"
	"x-bank-users/cerrors"
	"x-bank-users/ercodes"
)

type (
	Service struct {
		userStorageSeq int64
		userStorage    map[int64]storedUser
		userStorageMu  *sync.Mutex

		activationCodeCache   map[string]int64
		activationCodeCacheMu *sync.RWMutex
	}
)

func NewService() Service {
	return Service{
		userStorageSeq:        0,
		userStorage:           map[int64]storedUser{},
		userStorageMu:         &sync.Mutex{},
		activationCodeCache:   map[string]int64{},
		activationCodeCacheMu: &sync.RWMutex{},
	}
}

func (s *Service) CreateUser(_ context.Context, login, email string, passwordHash []byte) (int64, error) {
	s.userStorageMu.Lock()
	defer s.userStorageMu.Unlock()

	s.userStorageSeq++
	s.userStorage[s.userStorageSeq] = storedUser{
		Login:       login,
		Email:       email,
		Password:    passwordHash,
		IsActivated: false,
	}

	return s.userStorageSeq, nil
}

func (s *Service) ActivateUser(_ context.Context, userId int64) error {
	s.userStorageMu.Lock()
	defer s.userStorageMu.Unlock()

	user, ok := s.userStorage[userId]
	if !ok {
		return cerrors.NewErrorWithUserMessage(ercodes.UserNotFound, nil, "Пользователь не найден")
	}

	user.IsActivated = true
	s.userStorage[userId] = user

	return nil
}

func (s *Service) GenerateString(_ context.Context, set string, size int) (string, error) {
	return strings.Repeat(string([]rune(set)[0]), size), nil
}

func (s *Service) PutActivationCode(_ context.Context, code string, userId int64) error {
	s.activationCodeCacheMu.Lock()
	defer s.activationCodeCacheMu.Unlock()

	s.activationCodeCache[code] = userId
	return nil
}

func (s *Service) VerifyActivationCode(_ context.Context, code string) (int64, error) {
	s.activationCodeCacheMu.RLock()
	defer s.activationCodeCacheMu.RUnlock()

	userId, ok := s.activationCodeCache[code]
	if !ok {
		return 0, cerrors.NewErrorWithUserMessage(ercodes.ActivationCodeNotFound, nil, "Код активации не найден")
	}

	return userId, nil
}

func (s *Service) SendActivationCode(_ context.Context, email, code string) error {
	fmt.Printf("Письмо отправлено на %s: Ссылка на активации: https://example.com/?code=%s", email, code)
	return nil
}

func (s *Service) Hash(_ context.Context, b []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)
}
