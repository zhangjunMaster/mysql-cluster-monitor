package email

import (
	"strconv"
	"strings"

	gomail "gopkg.in/gomail.v2"
)

type Dialer struct {
	host     string
	port     int
	user     string
	password string
}

type Email struct {
	From    string
	To      []string
	Subject string
	Body    string
	Hosts   map[string]string
	Dialer  *Dialer
}

//var config, err = lib.ParseConfig()

func NewDialer(config map[string]map[string]string) (*Dialer, error) {
	dialerConfig := config["dialer"]
	port, err := strconv.Atoi(dialerConfig["port"])

	if err != nil {
		return nil, err
	}
	return &Dialer{
		dialerConfig["host"], port, dialerConfig["user"], dialerConfig["password"],
	}, nil
}

func NewEmail(config map[string]map[string]string) (*Email, error) {
	dialer, err := NewDialer(config)
	if err != nil {
		return nil, err
	}
	emailConfig := config["email"]
	receiver := strings.Split(emailConfig["to"], ",")
	return &Email{
		From:    emailConfig["from"],
		To:      receiver,
		Subject: emailConfig["subject"],
		Hosts:   config["hosts"],
		Dialer:  dialer,
	}, nil
}

func (e *Email) SendEmail() error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("To", e.To...)
	m.SetHeader("Subject", e.Subject)
	m.SetBody("text/html", e.Body)

	d := gomail.NewDialer(e.Dialer.host, e.Dialer.port, e.Dialer.user, e.Dialer.password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
