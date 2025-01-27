package mailer

import (
	"fmt"
	"net/smtp"
	"os"
)

// Функция для отправки письма
func SendEmail(to string, subject string, body string) error {
	from := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASSWORD")

	// Настроим SMTP клиент
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	// Настроим аутентификацию
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Формируем письмо
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body + "\r\n")

	// Отправляем письмо
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
