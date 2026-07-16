package auth

import (
	"context"
	"log/slog"
	"sso/server/internal/domain/models"
	"time"
)

type Auth struct {
	log         *slog.Logger //для работы с данными
	usrSaver    UserSaver
	usrProvider UserProvider
	appProvider AppProvider
	tokenTTL    time.Duration
}

type UserSaver interface {
	// User
	SaveUser(
		ctx context.Context,
		email string,
		passHash []byte,
	) (uid int64, err error)
}

// мб получать будем не только из бд
type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

type AppProvider interface {
	App(ctx context.Context, appID int) (models.App, error)
}

// New returns a new instance of AUTH service
func New(
	log *slog.Logger,
	UserSaver UserSaver,
	UserProvider UserProvider,
	AppProvider AppProvider,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		usrSaver:    UserSaver,
		usrProvider: UserProvider,
		appProvider: AppProvider,
		tokenTTL:    tokenTTL,
	}
}
