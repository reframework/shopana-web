package userStorageDto

import "github.com/google/uuid"

type CreateInput struct {
	Email       string `validate:"required,email"`
	FirstName   string `validate:"required,gt=0"`
	IsBlocked   *bool
	IsVerified  *bool
	Language    string  `validate:"required,gte=2,lte=5"`
	LastName    string  `validate:"required,gt=0"`
	Password    string  `validate:"required,gte=6"`
	PhoneNumber *string `validate:"omitempty,gte=10,lte=15"`
	Timezone    string  `validate:"required,timezone"`
	InvitedBy   int
}

type InviteInput struct {
	Email     string `validate:"required,email"`
	InvitedBy int    `validate:"required"`
}

type UpdateInput struct {
	ID          uuid.UUID `validate:"required"`
	Email       string    `validate:"omitempty,email"`
	FirstName   string    `validate:"omitempty,gt=0"`
	IsBlocked   *bool
	IsVerified  *bool
	Language    string  `validate:"omitempty,gte=2,lte=5"`
	LastName    string  `validate:"omitempty,gt=0"`
	Password    string  `validate:"omitempty,gte=6"`
	PhoneNumber *string `validate:"omitempty,gte=10,lte=15"`
	Timezone    string  `validate:"omitempty,timezone"`
}
