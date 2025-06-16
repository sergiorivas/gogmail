# GoGmail

A simple  Go library for sending emails through Gmail.

## Installation

```bash
go get github.com/sergiorivas/gogmail
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/sergiorivas/gogmail"
)

func main() {
	// Initialize client with your Gmail email and app password
	client := gogmail.New("your@gmail.com", "your-app-password")

	// Create a new email
	email := gogmail.NewEmail().
		AddTo("recipient@example.com").
		SetSubject("Hello from GoGmail").
		SetBody("This is a plain text message").
		SetHtmlBody("<h1>This is an HTML message</h1>").
		AddAttachment("file.pdf")

	// Send the email
	err := client.Send(email)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("Email sent successfully!")
}
```

### Features
- Simple initialization with Gmail credentials
- Support for multiple recipients, CC, and BCC
- Optional Reply-To address
- Both plain text and HTML email body
- File attachments support
- Fluent builder API for creating emails

### How to Get an App Password for Gmail
To use this library, you'll need to generate an app password for your Gmail account:

1. Go to your Google Account settings: https://myaccount.google.com/
2. Select "Security" from the left menu
3. Under "Signing in to Google," select "2-Step Verification" (must be enabled)
4. At the bottom of the page, select "App passwords"
5. Generate a new app password by selecting "Mail" as the app and "Other" as the device
6. Use the generated 16-character password in your application instead of your regular Gmail password

### Security Considerations
Never hardcode your app password in your source code. Use environment variables or a secure configuration manager.

### License
MIT
