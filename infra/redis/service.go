package redis

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
	"x-bank-users/cerrors"
	"x-bank-users/ercodes"
)

type (
	Service struct {
		db *redis.Client
	}
)

func NewService(password, host string, port int, database, maxCons int) (Service, error) {
	client := redis.NewClient(&redis.Options{
		Addr:           host + ":" + strconv.Itoa(port),
		Password:       password,
		DB:             database,
		MaxActiveConns: maxCons,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return Service{}, err
	}

	return Service{
		db: client,
	}, nil
}

func (s *Service) SaveActivationCode(ctx context.Context, code string, userId int64, ttl time.Duration) error {
	if err := s.db.Set(ctx, activationCodeKey+code, userId, ttl).Err(); err != nil {
		return s.wrapQueryError(err)
	}

	return nil
}

func (s *Service) VerifyActivationCode(ctx context.Context, code string) (int64, error) {
	userId, err := s.db.Get(ctx, activationCodeKey+code).Int64()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, cerrors.NewErrorWithUserMessage(ercodes.ActivationCodeNotFound, nil, "Код активации не найден")
		}
		return 0, s.wrapQueryError(err)
	}

	return userId, nil
}

func (s *Service) SaveRecoveryCode(ctx context.Context, code string, userId int64, ttl time.Duration) error {
	//TODO Реализовать
	panic("implement me")
}

func (s *Service) VerifyRecoveryCode(ctx context.Context, code string) (int64, error) {
	//TODO Реализовать
	panic("implement me")
}

func (s *Service) SaveRefreshToken(ctx context.Context, token string, userId int64, ttl time.Duration) error {
	//TODO Реализовать
	panic("implement me")
}

func (s *Service) VerifyRefreshToken(ctx context.Context, token string) (int64, error) {
	//TODO Реализовать
	panic("implement me")
}

func (s *Service) ExpireAllByUserId(ctx context.Context, userId int64) error {
	//TODO Реализовать
	panic("implement me")
}

func (s *Service) Save2FaCode(ctx context.Context, code string, userId int64, ttl time.Duration) error {
	//TODO Реализовать
	panic("implement me")
}

func (s *Service) Verify2FaCode(ctx context.Context, code string) (int64, error) {
	//TODO Реализовать
	panic("implement me")
}
