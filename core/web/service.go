package web

import "context"

type (
	Service struct {
		userStorage            UserStorage
		randomGenerator        RandomGenerator
		activationCodeCache    ActivationCodeCache
		activationCodeNotifier ActivationCodeNotifier
	}
)

func NewService(
	userStorage UserStorage,
	randomGenerator RandomGenerator,
	activationCodeCache ActivationCodeCache,
	activationCodeNotifier ActivationCodeNotifier,
) Service {
	return Service{
		userStorage:            userStorage,
		randomGenerator:        randomGenerator,
		activationCodeCache:    activationCodeCache,
		activationCodeNotifier: activationCodeNotifier,
	}
}

func (s *Service) SignUp() {
	// TODO Алёна
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
