package telegram

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"x-bank-users/cerrors"
	"x-bank-users/ercodes"
)

type (
	Service struct {
		client   *http.Client
		url      string
		login    string
		password string
	}
)

func NewService(URL, Login, Password string) Service {
	return Service{
		client:   &http.Client{},
		url:      URL,
		login:    Login,
		password: Password,
	}
}

func (s *Service) Send2FaCode(ctx context.Context, telegramId int64, code string) error {
	reqBody := strings.NewReader(fmt.Sprintf(`{"userId": %d, "code": "%s"}`, telegramId, code))

	req, err := http.NewRequestWithContext(ctx, "POST", s.url, reqBody)
	if err != nil {
		return cerrors.NewErrorWithUserMessage(ercodes.TelegramSendError, err, "Ошибка отправки кода")
	}

	req.SetBasicAuth(s.login, s.password)

	resp, err := s.client.Do(req)
	if err != nil {
		return cerrors.NewErrorWithUserMessage(ercodes.TelegramSendError, err, "Ошибка отправки кода")
	}

	if resp.StatusCode != http.StatusOK {
		return cerrors.NewErrorWithUserMessage(ercodes.TelegramSendError, nil, "Ошибка отправки кода")
	}

	return nil
}
