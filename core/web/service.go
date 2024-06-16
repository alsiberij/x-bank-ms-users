package web

import "context"

type (
	Service struct {
		userStorage            UserStorage
		randomGenerator        RandomGenerator
		activationCodeCache    ActivationCodeCache
		activationCodeNotifier ActivationCodeNotifier
		hashFunc               PasswordHasher
	}
)

func NewService(
	userStorage UserStorage,
	randomGenerator RandomGenerator,
	activationCodeCache ActivationCodeCache,
	activationCodeNotifier ActivationCodeNotifier,
	hashFunc PasswordHasher,
) Service {
	return Service{
		userStorage:            userStorage,
		randomGenerator:        randomGenerator,
		activationCodeCache:    activationCodeCache,
		activationCodeNotifier: activationCodeNotifier,
		hashFunc:               hashFunc,
	}
}

const (
	emailCodeLength  = 10
	emailCodeCharset = "0123456789"
	hashCost         = 10
)

func (s *Service) SignUp(ctx context.Context, login, password, email string) error {
	activationCode, err := s.randomGenerator.GenerateString(ctx, emailCodeCharset, emailCodeLength)

	if err != nil {
		return err
	}

	hash, err := s.hashFunc.HashPassword(ctx, []byte(password), hashCost)
	if err != nil {
		return err
	}

	userId, err := s.userStorage.CreateUser(ctx, login, email, hash)
	if err != nil {
		return err
	}

	err = s.activationCodeCache.PutActivationCode(ctx, activationCode, userId)
	if err != nil {
		return err
	}

	err = s.activationCodeNotifier.SendActivationCode(ctx, email, activationCode)
	return err
}

func (s *Service) ActivateAccount(ctx context.Context, code string) error {
	userId, err := s.activationCodeCache.VerifyActivationCode(ctx, code)
	if err != nil {
		return err
	}
	err = s.userStorage.ActivateUser(ctx, userId)
	if err != nil {
		return err
	}

	return nil
}
