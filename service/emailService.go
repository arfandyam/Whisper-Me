package service

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"

	"github.com/arfandyam/Whisper-Me/libs/exceptions"
	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/gin-gonic/gin"
)

type EmailService struct{}

func NewEmailService() EmailServiceInterface {
	return &EmailService{}
}

func (service *EmailService) SendEmailVerification(ctx *gin.Context, emailProperties *dto.EmailVerificationProperties) error {
	tmpl, err := template.ParseFiles(filepath.Join("templates", "emailVerification.html"))
	// tmpl, err := template.ParseFiles("../templates/emailVerification.html")
	if err != nil {
		err := exceptions.NewCustomError(http.StatusInternalServerError, "Failed to parse templates", err.Error())
		ctx.Error(err)
		return err
	}

	var renderedEmail bytes.Buffer
	if err := tmpl.Execute(&renderedEmail, emailProperties); err != nil {
		err := exceptions.NewCustomError(http.StatusInternalServerError, "Failed to write email", err.Error())
		ctx.Error(err)
		return err
	}

	go func(){
		auth := smtp.PlainAuth(
			"",
			os.Getenv("FROM_EMAIL"),
			os.Getenv("FROM_EMAIL_PASSWORD"),
			os.Getenv("GOOGLE_EMAIL_SMTP"),
		)
	
		headers := "MIME-version: 1.0;\nContent-Type: text/html; charset:\"UTF-8\";"
		message := "Subject: " + emailProperties.Subject + "\n" + headers + renderedEmail.String()
	
		if err := smtp.SendMail(
			os.Getenv("SMTP_ADDR"),
			auth,
			os.Getenv("FROM_EMAIL"),
			emailProperties.ToEmail,
			[]byte(message),
		); err != nil {
			err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to send email", err.Error())
			ctx.Error(err)
			log.Fatalf("Failed to send email to %v: %v", emailProperties.ToEmail, err)
		}
	}()

	return nil
}
