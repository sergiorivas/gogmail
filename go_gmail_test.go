package gogmail

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := New("test@gmail.com", "password")
	assert.NotNil(t, client)
	assert.Equal(t, "test@gmail.com", client.email)
	assert.Equal(t, "password", client.password)
}

func TestNewEmail(t *testing.T) {
	email := NewEmail()
	assert.NotNil(t, email)
	assert.Empty(t, email.To)
	assert.Empty(t, email.Subject)
}

func TestEmailBuilders(t *testing.T) {
	email := NewEmail().
		SetFrom("from@example.com").
		AddTo("to@example.com", "to2@example.com").
		AddCc("cc@example.com").
		AddBcc("bcc@example.com").
		SetReplyTo("reply@example.com").
		SetSubject("Test Subject").
		SetBody("Plain text body").
		SetHtmlBody("<h1>HTML Body</h1>").
		AddAttachment("test.txt")

	assert.Equal(t, "from@example.com", email.From)
	assert.Equal(t, []string{"to@example.com", "to2@example.com"}, email.To)
	assert.Equal(t, []string{"cc@example.com"}, email.Cc)
	assert.Equal(t, []string{"bcc@example.com"}, email.Bcc)
	assert.Equal(t, "reply@example.com", email.ReplyTo)
	assert.Equal(t, "Test Subject", email.Subject)
	assert.Equal(t, "Plain text body", email.PlainBody)
	assert.Equal(t, "<h1>HTML Body</h1>", email.HtmlBody)
	assert.Equal(t, []string{"test.txt"}, email.Attachments)
}

func TestSendIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	email := os.Getenv("TEST_EMAIL")
	password := os.Getenv("TEST_APP_PASSWORD")

	if email == "" || password == "" {
		t.Skip("TEST_EMAIL or TEST_APP_PASSWORD not set")
	}

	client := New(email, password)

	err := client.Send(NewEmail().
		AddTo(email).
		SetSubject("Test Email").
		SetBody("This is a test email").
		SetHtmlBody("<h1>This is a test email</h1>"))

	assert.NoError(t, err)
}
