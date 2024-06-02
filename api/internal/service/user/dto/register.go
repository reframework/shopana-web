package userServiceDto

type RegisterInput struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Timezone  string `json:"timezone" validate:"required"`
	Language  string `json:"language" validate:"required"`
}
