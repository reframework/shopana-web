package mailerService

import app "webapi/internal/app/interface"

type Service struct {
	app app.App
}

func New(app app.App) *Service {
	return &Service{
		app: app,
	}
}
