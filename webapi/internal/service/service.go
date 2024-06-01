package service

import (
	app "webapi/internal/app/interface"
	mailerService "webapi/internal/service/mailer"
	userService "webapi/internal/service/user"
)

type Service struct {
	user   *userService.Service
	mailer *mailerService.Service
}

func NewService(app app.App) *Service {
	return &Service{
		user:   userService.New(app),
		mailer: mailerService.New(app),
	}
}

func (s *Service) User() app.UserService {
	return s.user
}

func (s *Service) Mailer() app.MailerService {
	return s.mailer
}
