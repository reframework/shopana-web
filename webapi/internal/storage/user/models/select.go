package userStorageModel

import (
	"database/sql"
	"time"

	"webapi/internal/entity"

	"github.com/google/uuid"
)

type SelectModel struct {
	ID          string         `db:"id" goqu:"skipinsert,skipupdate"`
	CreatedAt   time.Time      `db:"created_at" goqu:"omitempty"`
	Email       string         `db:"email" goqu:"omitempty"`
	FirstName   string         `db:"first_name" goqu:"omitempty"`
	IsBlocked   bool           `db:"is_blocked" goqu:"omitempty"`
	IsVerified  bool           `db:"is_verified" goqu:"omitempty"`
	LastName    string         `db:"last_name" goqu:"omitempty"`
	Password    string         `db:"password" goqu:"omitempty"`
	PhoneNumber sql.NullString `db:"phone_number" goqu:"omitempty"`
	UpdatedAt   time.Time      `db:"updated_at" goqu:"omitempty"`
}

func (m *SelectModel) ToEntity() *entity.User {
	return &entity.User{
		ID:          uuid.MustParse(m.ID),
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		FirstName:   m.FirstName,
		LastName:    m.LastName,
		Email:       m.Email,
		Password:    m.Password,
		PhoneNumber: m.PhoneNumber.String,
		IsVerified:  m.IsVerified,
		IsBlocked:   m.IsBlocked,
	}
}
