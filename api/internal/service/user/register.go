package userService

import (
	"context"

	userServiceDto "webapi/internal/service/user/dto"
	userStorageDto "webapi/internal/storage/user/dto"
	appErrors "webapi/pkg/error"
	"webapi/pkg/jwt"

	"github.com/spf13/viper"
)

func (s *Service) Register(ctx context.Context, input *userServiceDto.RegisterInput) (*userServiceDto.Session, error) {
	if err := s.app.Validator().Struct(input); err != nil {
		return nil, appErrors.InvalidRequest.New(err)
	}

	isVerified := true
	id, err := s.app.Storage().User().Create(ctx, &userStorageDto.CreateInput{
		Email:      input.Email,
		FirstName:  input.FirstName,
		Language:   input.Language,
		LastName:   input.LastName,
		Password:   input.Password,
		Timezone:   input.Timezone,
		IsVerified: &isVerified,
	})
	if err != nil {
		return nil, err
	}

	user, err := s.app.Storage().User().FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	// defer s.SendVerifyEmail(ctx, user)

	jwtToken, err := jwt.CreateAuthToken(id, viper.GetString("AUTH_JWT_SECRET"))
	if err != nil {
		return nil, appErrors.Internal.New(err)
	}

	return &userServiceDto.Session{
		Jwt:  jwtToken,
		User: user,
	}, nil
}
