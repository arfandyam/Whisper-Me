package dto

type EmailVerificationProperties struct {
	ToEmail          []string
	Subject          string
	VerificationLink string
	IssuedAt         string
	ExpiredAt        string
}
