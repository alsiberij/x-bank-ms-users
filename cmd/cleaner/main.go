package main

import (
	"context"
	"x-bank-users/core/cleaner"
	"x-bank-users/infra/swissknife"
)

func main() {
	knife := swissknife.NewService()
	service := cleaner.NewService(&knife)
	if err := service.CleanExpiredUsers(context.Background()); err != nil {
		return
	}
}
