package userServiceDto

type VerifyEmailInput struct {
	Token string `json:"token"`
	Email string `json:"email"`
}
