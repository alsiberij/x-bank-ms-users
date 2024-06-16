package web

import "context"

type (
	UserStorage interface {
		CreateUser(ctx context.Context, login, email string, passwordHash []byte) (int64, error)
		ActivateUser(ctx context.Context, userId int64) error
	}

	RandomGenerator interface {
		GenerateString(ctx context.Context, set string, size int) (string, error)
	}

	ActivationCodeCache interface {
		PutActivationCode(ctx context.Context, code string, userId int64) error
		VerifyActivationCode(ctx context.Context, code string) (int64, error)
	}

	ActivationCodeNotifier interface {
		SendActivationCode(ctx context.Context, email, code string) error
	}

	PasswordHasher interface {
		HashPassword(ctx context.Context, b []byte, cost int) ([]byte, error)
	}
)
