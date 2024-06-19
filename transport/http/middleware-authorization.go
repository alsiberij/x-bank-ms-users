package http

// TODO Игорь
// Написать метод возвращающий мидлвару, которая будет проверять токен (заголовок Authorization).
// В случае успеха записывать *auth.Claims в контекст запроса. Если авторизоваться не удалось, возвращать ошибку.

func (t *Transport) authMiddleware(allow2Fa bool) middleware {
	panic("implement me")
}
