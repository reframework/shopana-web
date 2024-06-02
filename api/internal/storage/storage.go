package storage

import (
	app "webapi/internal/app/interface"
	configStorage "webapi/internal/storage/config"
	userStorage "webapi/internal/storage/user"
)

type Storage struct {
	user   app.UserStorage
	config app.ConfigStorage
}

func NewStorage(ap app.App) app.Storage {
	return &Storage{
		user:   userStorage.New(ap),
		config: configStorage.New(ap),
	}
}

func (r *Storage) User() app.UserStorage { return r.user }

func (r *Storage) Config() app.ConfigStorage { return r.config }
