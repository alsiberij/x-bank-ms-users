package swissknife

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
	"x-bank-users/cerrors"
	"x-bank-users/core/web"
	"x-bank-users/ercodes"
)

type (
	Service struct {
		userStorageSeq int64
		userStorage    map[int64]storedUser
		userStorageMu  *sync.Mutex

		strCodeCache   map[string]int64
		strCodeCacheMu *sync.RWMutex
	}
)

func NewService() Service {
	return Service{
		userStorageSeq: 0,
		userStorage:    map[int64]storedUser{},
		userStorageMu:  &sync.Mutex{},
		strCodeCache:   map[string]int64{},
		strCodeCacheMu: &sync.RWMutex{},
	}
}

func (s *Service) CreateUser(_ context.Context, login, email string, passwordHash []byte) (int64, error) {
	s.userStorageMu.Lock()
	defer s.userStorageMu.Unlock()

	s.userStorageSeq++
	s.userStorage[s.userStorageSeq] = storedUser{
		Login:           login,
		Email:           email,
		Password:        passwordHash,
		IsActivated:     false,
		HasPersonalData: false,
		TelegramId:      nil,
		CreatedAt:       time.Now(),
	}

	return s.userStorageSeq, nil
}

func (s *Service) GetSignInDataByLogin(_ context.Context, login string) (web.UserDataToSignIn, error) {
	s.userStorageMu.Lock()
	defer s.userStorageMu.Unlock()

	for id, user := range s.userStorage {
		if user.Login == login {
			return web.UserDataToSignIn{
				Id:              id,
				PasswordHash:    user.Password,
				TelegramId:      user.TelegramId,
				IsActivated:     user.IsActivated,
				HasPersonalData: user.HasPersonalData,
			}, nil
		}
	}

	return web.UserDataToSignIn{}, errors.New("user not found")
}

func (s *Service) GetSignInDataById(_ context.Context, id int64) (web.UserDataToSignIn, error) {
	s.userStorageMu.Lock()
	defer s.userStorageMu.Unlock()

	user, ok := s.userStorage[id]
	if !ok {
		return web.UserDataToSignIn{}, cerrors.NewErrorWithUserMessage(ercodes.UserNotFound, nil, "Пользователь не найден")
	}

	return web.UserDataToSignIn{
		Id:              id,
		PasswordHash:    user.Password,
		TelegramId:      user.TelegramId,
		IsActivated:     user.IsActivated,
		HasPersonalData: user.HasPersonalData,
	}, nil
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

func (s *Service) UserIdByLoginAndEmail(_ context.Context, login, email string) (int64, error) {
	s.userStorageMu.Lock()
	defer s.userStorageMu.Unlock()

	for id, user := range s.userStorage {
		if user.Login == login && user.Email == email {
			return id, nil
		}
	}

	return 0, errors.New("user not found")
}

func (s *Service) UpdatePassword(_ context.Context, id int64, passwordHash []byte) error {
	s.userStorageMu.Lock()
	defer s.userStorageMu.Unlock()

	user, ok := s.userStorage[id]
	if !ok {
		return cerrors.NewErrorWithUserMessage(ercodes.UserNotFound, nil, "Пользователь не найден")
	}

	user.Password = passwordHash
	s.userStorage[id] = user

	return nil
}

func (s *Service) SaveActivationCode(_ context.Context, code string, userId int64, _ time.Duration) error {
	s.strCodeCacheMu.Lock()
	defer s.strCodeCacheMu.Unlock()

	s.strCodeCache[code] = userId
	return nil
}

func (s *Service) VerifyActivationCode(_ context.Context, code string) (int64, error) {
	s.strCodeCacheMu.RLock()
	defer s.strCodeCacheMu.RUnlock()

	userId, ok := s.strCodeCache[code]
	if !ok {
		return 0, cerrors.NewErrorWithUserMessage(ercodes.ActivationCodeNotFound, nil, "Код активации не найден")
	}

	return userId, nil
}

func (s *Service) SendActivationCode(_ context.Context, email, code string) error {
	fmt.Printf("Письмо отправлено на %s: Ссылка на активации: https://example.com/?code=%s\n", email, code)
	return nil
}

func (s *Service) SendRecoveryCode(_ context.Context, email, code string) error {
	fmt.Printf("Письмо отправлено на %s: Код восстановления %s", email, code)
	return nil
}

func (s *Service) SaveRefreshToken(ctx context.Context, token string, userId int64, _ time.Duration) error {
	s.strCodeCacheMu.Lock()
	defer s.strCodeCacheMu.Unlock()

	s.strCodeCache[token] = userId
	return nil
}

func (s *Service) VerifyRefreshToken(_ context.Context, token string) (int64, error) {
	s.strCodeCacheMu.RLock()
	defer s.strCodeCacheMu.RUnlock()

	userId, ok := s.strCodeCache[token]
	if !ok {
		return 0, cerrors.NewErrorWithUserMessage(ercodes.ActivationCodeNotFound, nil, "Код активации не найден")
	}

	return userId, nil
}

func (s *Service) Save2FaCode(_ context.Context, code string, userId int64, _ time.Duration) error {
	s.strCodeCacheMu.Lock()
	defer s.strCodeCacheMu.Unlock()

	s.strCodeCache[code] = userId
	return nil
}

func (s *Service) Verify2FaCode(_ context.Context, code string) (int64, error) {
	s.strCodeCacheMu.RLock()
	defer s.strCodeCacheMu.RUnlock()

	userId, ok := s.strCodeCache[code]
	if !ok {
		return 0, cerrors.NewErrorWithUserMessage(ercodes.ActivationCodeNotFound, nil, "Код активации не найден")
	}

	return userId, nil
}

func (s *Service) Send2FaCode(_ context.Context, telegramId int64, code string) error {
	fmt.Printf("Отправлено в телеграм id %d: Код %s\n", telegramId, code)
	return nil
}

func (s *Service) SaveRecoveryCode(_ context.Context, code string, userId int64, ttl time.Duration) error {
	s.strCodeCacheMu.Lock()
	defer s.strCodeCacheMu.Unlock()

	s.strCodeCache[code] = userId
	return nil
}

func (s *Service) VerifyRecoveryCode(_ context.Context, code string) (int64, error) {
	s.strCodeCacheMu.RLock()
	defer s.strCodeCacheMu.RUnlock()

	userId, ok := s.strCodeCache[code]
	if !ok {
		return 0, cerrors.NewErrorWithUserMessage(ercodes.ActivationCodeNotFound, nil, "Код активации не найден")
	}

	return userId, nil
}

func (s *Service) ExpireAllByUserId(_ context.Context, userId int64) error {
	s.strCodeCacheMu.Lock()
	defer s.strCodeCacheMu.Unlock()

	for k, v := range s.strCodeCache {
		if v == userId {
			delete(s.strCodeCache, k)
		}
	}
	fmt.Println(s.strCodeCache)
	return nil
}

func (s *Service) DeleteUsersWithExpiredActivation(ctx context.Context, expirationTime time.Duration) error {
	s.userStorageMu.Lock()
	defer s.userStorageMu.Unlock()

	timeNow := time.Now()

	for key, value := range s.userStorage {
		if !value.IsActivated && value.CreatedAt.Sub(timeNow) >= expirationTime {
			if err := s.ExpireAllByUserId(ctx, key); err != nil {
				return err
			}
			delete(s.userStorage, key)
		}
	}
	return nil
}
