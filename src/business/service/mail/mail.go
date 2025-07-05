package mail

import (
	"context"
	"fmt"

	"github.com/NupalHariz/DD/src/utils/config"
	"github.com/reyhanmichiels/go-pkg/v2/codes"
	"github.com/reyhanmichiels/go-pkg/v2/errors"
	"github.com/reyhanmichiels/go-pkg/v2/log"
	"gopkg.in/gomail.v2"
)

type Interface interface {
	SendEmail(ctx context.Context, email string, subject string, body string) error
}

type mail struct {
	log log.Interface
	cfg config.Mail
}

type InitParam struct {
	Log log.Interface
	Cfg config.Mail
}

func Init(param InitParam) Interface {
	return &mail{
		log: param.Log,
		cfg: param.Cfg,
	}
}

func (g *mail) SendEmail(ctx context.Context, email string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", g.cfg.Username)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	g.log.Debug(ctx, fmt.Sprintf("sending email to %s with body %s", email, body))

	d := gomail.NewDialer(g.cfg.Host, int(g.cfg.Port), g.cfg.Username, g.cfg.Password)

	if err := d.DialAndSend(m); err != nil {
		g.log.Error(ctx, fmt.Sprintf("failed to send email to %s, error: %v", email, err))
		return  errors.NewWithCode(codes.CodeInternalServerError, err.Error())
	}

	g.log.Debug(ctx, fmt.Sprintf("success sending email to %s with body %s", email, body))

	return nil
}