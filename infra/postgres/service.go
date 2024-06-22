package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"time"
	"x-bank-users/cerrors"
	"x-bank-users/core/web"
	"x-bank-users/ercodes"
)

const (
	uniqueLoginConstraint = `users_login_key`
	uniqueEmailConstraint = `users_email_key`
)

type (
	Service struct {
		db *sql.DB
	}
)

func NewService(login, password, host string, port int, database string, maxCons int) (Service, error) {
	db, err := sql.Open("pgx", fmt.Sprintf("postgres://%s:%s@%s:%d/%s", login, password, host, port, database))
	if err != nil {
		return Service{}, err
	}

	db.SetMaxOpenConns(maxCons)

	if err = db.Ping(); err != nil {
		return Service{}, err
	}

	return Service{
		db: db,
	}, err
}

func (s *Service) CreateUser(ctx context.Context, login, email string, passwordHash []byte) (int64, error) {
	const query = `INSERT INTO users (login, email, password) VALUES (@login, @email, @password) RETURNING id`

	row := s.db.QueryRowContext(ctx, query,
		pgx.NamedArgs{
			"login":    login,
			"email":    email,
			"password": passwordHash,
		},
	)

	if err := row.Err(); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.ConstraintName {
			case uniqueLoginConstraint:
				return 0, cerrors.NewErrorWithUserMessage(ercodes.LoginAlreadyTaken, nil, "Логин уже занят")
			case uniqueEmailConstraint:
				return 0, cerrors.NewErrorWithUserMessage(ercodes.EmailAlreadyTaken, nil, "Емейл уже занят")
			}
		}
		return 0, s.wrapQueryError(err)
	}

	var userId int64
	if err := row.Scan(&userId); err != nil {
		return 0, s.wrapScanError(err)
	}

	return userId, nil
}

func (s *Service) GetSignInDataByLogin(ctx context.Context, login string) (web.UserDataToSignIn, error) {
	//TODO Реализовать
	panic("implement me")
}

func (s *Service) GetSignInDataById(ctx context.Context, id int64) (web.UserDataToSignIn, error) {
	//TODO Реализовать
	panic("implement me")
}

func (s *Service) ActivateUser(ctx context.Context, userId int64) error {
	//TODO Реализовать
	panic("implement me")
}

func (s *Service) UserIdByLoginAndEmail(ctx context.Context, login, email string) (int64, error) {
	//TODO Реализовать
	panic("implement me")
}

func (s *Service) UpdatePassword(ctx context.Context, id int64, passwordHash []byte) error {
	//TODO Реализовать
	panic("implement me")
}

func (s *Service) UpdateTelegramId(ctx context.Context, telegramId *int64, userId int64) error {
	//TODO Реализовать
	panic("implement me")
}

func (s *Service) GetUserDataById(ctx context.Context, userId int64) (*web.UserPersonalData, error) {
	//TODO Реализовать
	panic("implement me")
}

func (s *Service) DeleteUsersWithExpiredActivation(ctx context.Context, expirationTime time.Duration) error {
	//TODO Реализовать
	panic("implement me")
}
