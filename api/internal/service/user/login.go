package userService

import (
	"context"

	userServiceDto "webapi/internal/service/user/dto"
	appErrors "webapi/pkg/error"
	"webapi/pkg/helpers"
	"webapi/pkg/jwt"

	"github.com/spf13/viper"
)

func (s *Service) Login(ctx context.Context, input *userServiceDto.LoginInput) (*userServiceDto.Session, error) {
	if err := s.app.Validator().Struct(input); err != nil {
		return nil, appErrors.InvalidRequest.New(err)
	}

	user, err := s.app.Storage().User().FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	err = helpers.VerifyHashedPassword(input.Password, user.Password)
	if err != nil {
		return nil, appErrors.AccessDenied.New(err)
	}

	jwtToken, err := jwt.CreateAuthToken(user.ID, viper.GetString("AUTH_JWT_SECRET"))
	if err != nil {
		return nil, appErrors.Internal.New(err)
	}

	return &userServiceDto.Session{
		Jwt:  jwtToken,
		User: user,
	}, nil
}
