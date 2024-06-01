package mailerServiceDto

type SendEmailInput struct {
	To   string `json:"to" validate:"required"`
	Data any    `json:"data" validate:"required"`
}
