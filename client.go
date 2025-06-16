package gogmail

import (
	"crypto/tls"

	gomail "gopkg.in/mail.v2"
)

type Client struct {
	email    string
	password string
	dialer   *gomail.Dialer
}

func New(email, password string) *Client {
	dialer := gomail.NewDialer("smtp.gmail.com", 587, email, password)
	dialer.TLSConfig = &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         "smtp.gmail.com",
	}

	return &Client{
		email:    email,
		password: password,
		dialer:   dialer,
	}
}

func (c *Client) Send(email *Email) error {
	m := gomail.NewMessage()

	from := email.From
	if from == "" {
		from = c.email
	}

	m.SetHeader("From", from)
	m.SetHeader("To", email.To...)
	m.SetHeader("Subject", email.Subject)

	if len(email.Cc) > 0 {
		m.SetHeader("Cc", email.Cc...)
	}

	if len(email.Bcc) > 0 {
		m.SetHeader("Bcc", email.Bcc...)
	}

	if email.ReplyTo != "" {
		m.SetHeader("Reply-To", email.ReplyTo)
	}

	// If both plain and HTML body are provided, use both with HTML as alternative
	if email.PlainBody != "" && email.HtmlBody != "" {
		m.SetBody("text/plain", email.PlainBody)
		m.AddAlternative("text/html", email.HtmlBody)
	} else if email.HtmlBody != "" {
		// If only HTML body is provided, send only HTML
		m.SetBody("text/html", email.HtmlBody)
	} else {
		// Default to plain text if only plain body is provided or as fallback
		m.SetBody("text/plain", email.PlainBody)
	}

	for _, attachment := range email.Attachments {
		m.Attach(attachment)
	}

	return c.dialer.DialAndSend(m)
}
