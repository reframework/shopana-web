package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	FirstName   string
	LastName    string
	Email       string
	Password    string
	PhoneNumber string
	IsVerified  bool
	IsBlocked   bool
}

type TenantInvite struct {
	ID        int
	Email     string
	InvitedBy int
	InvitedAt time.Time
}
