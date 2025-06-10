package libs

import (
	"bytes"
	"html/template"
	"math"
	"math/rand"
	"strings"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
		if ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || ('0' <= char && char <= '9') || char == ' ' || char == '-' {
			slugTrim.WriteByte(char)
		}
	}

	slug = slugTrim.String() + "-" + itemId.String()[:13]
	return slug
}

func CalculateOffset(page int, fetchPerPage int) int {
	return (page - 1) * fetchPerPage
}

func SlugToBase62(slug string) string {
	numericString := 1
	for _, char := range slug {
		if int(char) >= 48 && int(char) <= 57 {
			numericString += (int(char) - 48)
		} else if int(char) >= 65 && int(char) <= 90 {
			numericString += (int(char) - 65 + 11)
		} else if int(char) >= 97 && int(char) <= 122 {
			numericString += (int(char) - 97 + 37)
		} else if int(char) == 45 {
		    numericString += (int(char) - 45 + 63)
		}
	}

	salt := int(math.Ceil(rand.Float64()*1500)*50*25)
	numericString *= salt

	return Base62Encode(numericString)
}

func Base62Encode(num int) string {
	base62Chars := "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base := 62
	result := ""

	for num > 0 {
		var remainder int
		num = num / base
		remainder = num % base
		result = string(base62Chars[remainder]) + result
	}

	return result
}