package configStorage

import app "webapi/internal/app/interface"

type Storage struct {
	app app.App
}

func New(app app.App) *Storage {
	return &Storage{
		app: app,
	}
}
