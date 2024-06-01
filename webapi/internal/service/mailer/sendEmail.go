package mailerService

import (
	"context"
	"crypto/tls"

	mailerServiceDto "webapi/internal/service/mailer/dto"
	appErrors "webapi/pkg/error"

	"github.com/aymerick/raymond"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func (r *Service) SendEmail(ctx context.Context, input *mailerServiceDto.SendEmailInput) error {
	template := ""

	m := gomail.NewMessage()

	m.SetHeader("To", input.To)
	m.SetHeader("From", viper.GetString("API_EMAIL_FROM"))
	m.SetHeader("Reply-To", viper.GetString("API_EMAIL_REPLY_TO"))
	m.SetHeader("Subject", "Subject")

	tpl, err := raymond.Parse(template)
	if err != nil {
		return appErrors.Internal.New(err, "failed to parse template")
	}

	body, err := tpl.Exec(input.Data)
	if err != nil {
		return appErrors.Internal.New(err, "failed to execute template")
	}

	m.SetBody("text/html", body)

	d := gomail.NewDialer(
		viper.GetString("API_EMAIL_SMTP_HOST"),
		viper.GetInt("API_EMAIL_SMTP_PORT"),
		viper.GetString("API_EMAIL_SMTP_USER"),
		viper.GetString("API_EMAIL_SMTP_PASSWORD"),
	)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return appErrors.Internal.New(err, "failed to send email")
	}

	return nil
}
