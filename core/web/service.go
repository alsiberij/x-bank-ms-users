package web

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

}

func (s *Service) VerifyAccount() {

}
