package libs

import (
	"bytes"
	"html/template"
	"strings"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateEmailBody(verificationLink string) (string, error) {
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

func ToSlug(itemString string, itemId uuid.UUID) string {

	// lowercase and replace whitespace and underscore to hyphens
	slug := strings.ToLower(itemString)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "_", "-")

	// Remove non-alphanumeric
	var slugTrim strings.Builder
	for i := 0; i < len(slug); i++ {
		char := slug[i]
		if ('a' <= char && char <= 'z') ||('A' <= char && char <= 'Z') ||('0' <= char && char <= '9') || char == ' ' || char == '-'{
			slugTrim.WriteByte(char)
		}
	}

	slug = slugTrim.String() + "-" + itemId.String()[:13]

	return slug
}

func CalculateOffset(page int, fetchPerPage int) int {
	return (page-1) * fetchPerPage
}
