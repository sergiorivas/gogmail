package gogmail

type Email struct {
	From        string
	To          []string
	Cc          []string
	Bcc         []string
	ReplyTo     string
	Subject     string
	PlainBody   string
	HtmlBody    string
	Attachments []string
}

func NewEmail() *Email {
	return &Email{
		To:          []string{},
		Cc:          []string{},
		Bcc:         []string{},
		Attachments: []string{},
	}
}

func (e *Email) SetFrom(from string) *Email {
	e.From = from
	return e
}

func (e *Email) AddTo(to ...string) *Email {
	e.To = append(e.To, to...)
	return e
}

func (e *Email) AddCc(cc ...string) *Email {
	e.Cc = append(e.Cc, cc...)
	return e
}

func (e *Email) AddBcc(bcc ...string) *Email {
	e.Bcc = append(e.Bcc, bcc...)
	return e
}

func (e *Email) SetReplyTo(replyTo string) *Email {
	e.ReplyTo = replyTo
	return e
}

func (e *Email) SetSubject(subject string) *Email {
	e.Subject = subject
	return e
}

func (e *Email) SetBody(body string) *Email {
	e.PlainBody = body
	return e
}

func (e *Email) SetHtmlBody(htmlBody string) *Email {
	e.HtmlBody = htmlBody
	return e
}

func (e *Email) AddAttachment(filepath ...string) *Email {
	e.Attachments = append(e.Attachments, filepath...)
	return e
}
