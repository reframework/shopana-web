package app

import (
	"context"

	"webapi/internal/entity"
	userStorageDto "webapi/internal/storage/user/dto"

	"github.com/google/uuid"
)

type Storage interface {
	User() UserStorage
	Config() ConfigStorage
}

type UserStorage interface {
	Create(ctx context.Context, input *userStorageDto.CreateInput) (uuid.UUID, error)
	FindByPhone(ctx context.Context, phone string) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindById(ctx context.Context, id uuid.UUID) (*entity.User, error)
}

type ConfigStorage interface {
	IsReady(ctx context.Context) (bool, error)
	SetReady(ctx context.Context) error
}
