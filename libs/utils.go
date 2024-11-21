package libs

import (
	"bytes"
	"html/template"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateEmailBody(verificationLink string) (string, error){
	tmpl, err := template.ParseFiles("../templates/emailVerification.html")
	if err != nil {
		return "", err
	}

	var rendered bytes.Buffer
	if err := tmpl.Execute(&rendered, verificationLink); err != nil {
		return "", err
	}

	return rendered.String(), nil
}

