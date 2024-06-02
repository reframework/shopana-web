package userService

import (
	"context"

	ctxutil "webapi/internal/app/ctx"
	"webapi/internal/entity"
	appErrors "webapi/pkg/error"
	"webapi/pkg/jwt"

	"github.com/spf13/viper"
)

func (s *Service) Me(ctx context.Context, token string) (*entity.User, error) {
	if user, ok := ctxutil.User(ctx); ok {
		return user, nil
	}

	userId, err := jwt.ValidateAuthToken(token, viper.GetString("API_JWT_SECRET"))
	if err != nil {
		return nil, appErrors.AccessDenied.New(err)
	}

	user, err := s.app.Storage().User().FindById(ctx, userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}
