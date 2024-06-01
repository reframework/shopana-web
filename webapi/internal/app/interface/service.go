package app

import (
	"context"

	"webapi/internal/entity"
	mailerServiceDto "webapi/internal/service/mailer/dto"
	userServiceDto "webapi/internal/service/user/dto"
)

type Service interface {
	User() UserService
	Mailer() MailerService
}

type UserService interface {
	Me(ctx context.Context, token string) (*entity.User, error)
	Register(ctx context.Context, input *userServiceDto.RegisterInput) (
		*userServiceDto.Session, error)
	Login(ctx context.Context, input *userServiceDto.LoginInput) (*userServiceDto.Session, error)
}

type MailerService interface {
	SendEmail(context.Context, *mailerServiceDto.SendEmailInput) error
}
