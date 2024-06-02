package configStorage

import (
	"context"
	"database/sql"
	"errors"

	"webapi/internal/storage/table"
	appErrors "webapi/pkg/error"

	"github.com/doug-martin/goqu/v9"
)

func (r *Storage) IsReady(ctx context.Context) (bool, error) {
	query, args, err := goqu.
		Select("initialized").
		From(goqu.T(table.Config)).
		Where(goqu.Ex{"initialized": true}).
		ToSQL()
	if err != nil {
		return false, appErrors.Internal.New(err)
	}

	var result bool
	err = r.app.DB().Get(&result, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, appErrors.Internal.New(err)
	}

	return result, nil
}

func (r *Storage) SetReady(ctx context.Context) error {
	query, args, err := goqu.
		Insert(goqu.T(table.Config)).
		Rows(goqu.Record{
			"initialized": true,
		}).
		ToSQL()
	if err != nil {
		return appErrors.Internal.New(err)
	}

	_, err = r.app.DB().Exec(query, args...)
	if err != nil {
		return appErrors.Internal.New(err)
	}

	return nil
}
