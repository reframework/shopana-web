package trx

import (
	isql "webapi/pkg/postgres/interface"

	"github.com/google/uuid"
)

func ExecVal[T any](db isql.Runner, fn func(tx isql.BaseRunner) (T, error)) (T, error) {
	tx, err := db.Beginx()
	if err != nil {
		var zero T
		return zero, err
	}

	defer tx.Rollback()

	result, err := fn(tx)
	if err != nil {
		return result, err
	}

	err = tx.Commit()
	if err != nil {
		var zero T
		return zero, err
	}

	return result, nil
}

func Exec(db isql.Runner, fn func(tx isql.BaseRunner) error) error {
	_, err := ExecVal(db, func(tx isql.BaseRunner) (any, error) {
		return nil, fn(tx)
	})

	return err
}

func ExecReturningId(tx isql.BaseRunner, query string, args ...any) (int, error) {
	stmt, err := tx.Prepare(query)
	if err != nil {
		return 0, err
	}

	var id int
	if err := stmt.QueryRow(args...).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func ExecReturningUUID(tx isql.BaseRunner, query string, args ...any) (uuid.UUID, error) {
	stmt, err := tx.Prepare(query)
	if err != nil {
		return uuid.Nil, err
	}

	var id string
	if err := stmt.QueryRow(args...).Scan(&id); err != nil {
		return uuid.Nil, err
	}

	parsed, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, err
	}

	return parsed, nil
}
