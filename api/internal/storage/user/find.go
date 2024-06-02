package userStorage

import (
	"context"
	"database/sql"

	"webapi/internal/entity"
	"webapi/internal/storage/table"
	userStorageModel "webapi/internal/storage/user/models"
	appErrors "webapi/pkg/error"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

func (r *Storage) FindByPhone(ctx context.Context, phone string) (*entity.User, error) {
	return r.findUserBy(ctx, findUserByInput{Phone: phone})
}

func (r *Storage) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	return r.findUserBy(ctx, findUserByInput{Email: email})
}

func (r *Storage) FindById(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	return r.findUserBy(ctx, findUserByInput{ID: id})
}

type findUserByInput struct {
	ID    uuid.UUID
	Email string
	Phone string
}

func (r *Storage) findUserBy(_ context.Context, input findUserByInput) (*entity.User, error) {
	sd := goqu.
		Dialect("postgres").
		Select(
			goqu.C("id"),
			goqu.C("email"),
			goqu.C("last_name"),
			goqu.C("first_name"),
			goqu.C("is_blocked"),
			goqu.C("is_verified"),
			goqu.C("password"),
			goqu.C("phone_number"),
		).
		From(goqu.T(table.User))

	if input.ID != uuid.Nil {
		sd = sd.Where(goqu.Ex{"id": input.ID})
	} else if input.Email != "" {
		sd = sd.Where(goqu.Ex{"email": input.Email})
	} else if input.Phone != "" {
		sd = sd.Where(goqu.Ex{"phone": input.Phone})
	} else {
		return nil, appErrors.Internal.New(nil, "id, email or phone is required")
	}

	query, args, err := sd.ToSQL()
	if err != nil {
		return nil, appErrors.Internal.New(err)
	}

	userModel := userStorageModel.SelectModel{}
	err = r.app.DB().Get(&userModel, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, appErrors.NotFound.New(err)
		}

		return nil, appErrors.Internal.New(err)
	}

	return userModel.ToEntity(), nil
}
