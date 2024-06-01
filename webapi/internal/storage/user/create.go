package userStorage

import (
	"context"

	"webapi/internal/storage/table"
	"webapi/pkg/helpers"
	trx "webapi/pkg/postgres/transaction"

	appErrors "webapi/pkg/error"

	userStorageDto "webapi/internal/storage/user/dto"
	userStorageModel "webapi/internal/storage/user/models"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

func (r *Storage) Create(ctx context.Context, input *userStorageDto.CreateInput) (uuid.UUID, error) {
	err := r.app.Validator().Struct(input)
	if err != nil {
		return uuid.Nil, appErrors.Internal.New(err)
	}

	hashedPassword, err := helpers.HashPassword(input.Password)
	if err != nil {
		return uuid.Nil, appErrors.Internal.New(err)
	}

	query, args, err := goqu.Dialect("postgres").
		Insert(goqu.T(table.User)).
		Rows(userStorageModel.MutateModel{
			FirstName:   input.FirstName,
			LastName:    input.LastName,
			Email:       input.Email,
			Password:    hashedPassword,
			PhoneNumber: input.PhoneNumber,
			IsBlocked:   input.IsBlocked,
			IsVerified:  input.IsVerified,
			Timezone:    input.Timezone,
			Language:    input.Language,
		}).
		Returning("id").
		ToSQL()
	if err != nil {
		return uuid.Nil, appErrors.Internal.New(err)
	}

	id, err := trx.ExecReturningUUID(r.app.DB(), query, args...)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
