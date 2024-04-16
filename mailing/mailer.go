package mailing

import (
	"dsr-mail-service/types"
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(data types.EmailRequest) {
	fmt.Println(data)
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASSWORD")

	to := []string{
		data.ToEmail,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "Subject: Test email from Go!\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := fmt.Sprintf(`
		<html>
			<body>
				<h1>You have a new email!</h1>
				<h4>From: %s %s</h4>
				<h4>Email: %s</h4>
				<h4>Message: %s</h4>
			</body>
		</html>`, data.FirstName, data.LastName, data.ToEmail, data.Message)

	message := []byte(subject + mime + body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	fmt.Println("Sending email...")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email Sent Successfully!")
}
