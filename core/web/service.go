package web

import (
	"context"
	"github.com/google/uuid"
	"time"
	"x-bank-users/auth"
	"x-bank-users/cerrors"
	"x-bank-users/ercodes"
)

type (
	Service struct {
		userStorage            UserStorage
		randomGenerator        RandomGenerator
		activationCodeCache    ActivationCodeStorage
		activationCodeNotifier AuthNotifier
		passwordHasher         PasswordHasher
		refreshTokenStorage    RefreshTokenStorage
		twoFactorCodeStorage   TwoFactorCodeStorage
		twoFactorCodeNotifier  TwoFactorCodeNotifier
		recoveryCodeStorage    RecoveryCodeStorage
	}
)

func NewService(
	userStorage UserStorage,
	randomGenerator RandomGenerator,
	activationCodeCache ActivationCodeStorage,
	activationCodeNotifier AuthNotifier,
	passwordHasher PasswordHasher,
	refreshTokenStorage RefreshTokenStorage,
	twoFactorCodeStorage TwoFactorCodeStorage,
	twoFactorCodeNotifier TwoFactorCodeNotifier,
	recoveryCodeStorage RecoveryCodeStorage,
) Service {
	return Service{
		userStorage:            userStorage,
		randomGenerator:        randomGenerator,
		activationCodeCache:    activationCodeCache,
		activationCodeNotifier: activationCodeNotifier,
		passwordHasher:         passwordHasher,
		refreshTokenStorage:    refreshTokenStorage,
		twoFactorCodeStorage:   twoFactorCodeStorage,
		twoFactorCodeNotifier:  twoFactorCodeNotifier,
		recoveryCodeStorage:    recoveryCodeStorage,
	}
}

const (
	emailCodeLength  = 10
	emailCodeCharset = "0123456789"
	emailCodeTtl     = time.Hour * 24

	hashCost = 10

	claimsTtl = time.Minute * 5

	refreshTokenCharset = ".-"
	refreshTokenSize    = 2048
	refreshTokenTtl     = time.Hour * 24 * 7

	twoFactorCodeCharset = "0123456789"
	twoFactorCodeSize    = 6
	TwoFactorCodeTtl     = time.Minute * 5

	recoveryCodeCharset = "ij"
	recoveryCodeSize    = 16
	recoveryCodeTtl     = time.Minute * 5
)

func (s *Service) SignUp(ctx context.Context, login, password, email string) error {
	activationCode, err := s.randomGenerator.GenerateString(ctx, emailCodeCharset, emailCodeLength)

	if err != nil {
		return err
	}

	hash, err := s.passwordHasher.HashPassword(ctx, []byte(password), hashCost)
	if err != nil {
		return err
	}

	userId, err := s.userStorage.CreateUser(ctx, login, email, hash)
	if err != nil {
		return err
	}

	err = s.activationCodeCache.SaveActivationCode(ctx, activationCode, userId, emailCodeTtl)
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

func (s *Service) SignIn(ctx context.Context, login, password string) (SignInResult, error) {
	userData, err := s.userStorage.GetSignInDataByLogin(ctx, login)
	if err != nil {
		return SignInResult{}, err
	}

	err = s.passwordHasher.CompareHashAndPassword(ctx, password, userData.PasswordHash)
	if err != nil {
		return SignInResult{}, err
	}

	if !userData.IsActivated {
		return SignInResult{}, cerrors.NewErrorWithUserMessage(ercodes.AccountNotActivated, nil, "Аккаунт не активирован")
	}

	var refreshToken string
	if userData.TelegramId != nil {
		refreshToken, err = s.randomGenerator.GenerateString(ctx, refreshTokenCharset, refreshTokenSize)
		if err != nil {
			return SignInResult{}, err
		}
		if err = s.refreshTokenStorage.SaveRefreshToken(ctx, refreshToken, userData.Id, refreshTokenTtl); err != nil {
			return SignInResult{}, err
		}
	} else {
		twoFactorCode, err := s.randomGenerator.GenerateString(ctx, twoFactorCodeCharset, twoFactorCodeSize)
		if err != nil {
			return SignInResult{}, err
		}
		if err = s.twoFactorCodeStorage.Save2FaCode(ctx, twoFactorCode, userData.Id, TwoFactorCodeTtl); err != nil {
			return SignInResult{}, err
		}
		if err = s.twoFactorCodeNotifier.Send2FaCode(ctx, *userData.TelegramId, twoFactorCode); err != nil {
			return SignInResult{}, err
		}
	}
	date := time.Now()

	claims := auth.Claims{
		Id:              uuid.New().String(),
		IssuedAt:        date.Unix(),
		ExpiresAt:       date.Add(claimsTtl).Unix(),
		Sub:             userData.Id,
		Is2FAToken:      userData.TelegramId != nil,
		HasPersonalData: userData.HasPersonalData,
	}

	return SignInResult{
		AccessClaims: claims,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Service) SignIn2FA(ctx context.Context, claims auth.Claims, code string) (SignInResult, error) {
	// TODO Игорь
	// 1. Проверка 2FA кода, извлечение userId
	// 2. Сравнение userId из claims, должны совпадать
	// 3. Поиск юзера по id
	// 4. Генерируем и сохраняем рефреш токен
	// 5. Формируем auth.Claims

	return SignInResult{}, nil
}

func (s *Service) Recovery(ctx context.Context, login, email string) error {
	userId, err := s.userStorage.UserIdByLoginAndEmail(ctx, login, email)
	if err != nil {
		return err
	}

	recoveryCode, err := s.randomGenerator.GenerateString(ctx, recoveryCodeCharset, recoveryCodeSize)
	if err != nil {
		return err
	}

	err = s.recoveryCodeStorage.SaveRecoveryCode(ctx, recoveryCode, userId, recoveryCodeTtl)
	if err != nil {
		return err
	}

	err = s.activationCodeNotifier.SendRecoveryCode(ctx, email, recoveryCode)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) RecoveryCode(ctx context.Context, code, password string) error {
	userId, err := s.recoveryCodeStorage.VerifyRecoveryCode(ctx, code)
	if err != nil {
		return err
	}

	hashedPassword, err := s.passwordHasher.HashPassword(ctx, []byte(password), hashCost)
	if err != nil {
		return err
	}

	err = s.userStorage.UpdatePassword(ctx, userId, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}
