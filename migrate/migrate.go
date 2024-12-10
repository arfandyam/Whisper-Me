package main

import (
	"log"

	"github.com/arfandyam/Whisper-Me/initializers"
	"github.com/arfandyam/Whisper-Me/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnDB()
}

func main() {
	/*Pisah AutoMigrate untuk tabel yg punya constraint karena dalam prosesnya tabel
	terkadang belum terbuat tapi sudah ada constraint terhadap tabel lain*/
	initializers.DB.AutoMigrate(&models.User{}, &models.Session{})
	initializers.DB.AutoMigrate(&models.Question{})
	initializers.DB.AutoMigrate(&models.Response{})

	// Persiapan Full Text Search untuk table questions
	if err := initializers.DB.Exec("ALTER TABLE questions ADD COLUMN question_vector tsvector").Error; err != nil {
		log.Fatal("err:", err)
	}
	if err := initializers.DB.Exec("UPDATE questions SET question_vector = to_tsvector('indonesian', topic || ' ' || question)").Error; err != nil {
		log.Fatal("err:", err)
	}
	if err := initializers.DB.Exec("CREATE INDEX idx_question_vector ON questions USING gin(question_vector)").Error; err != nil {
		log.Fatal("err:", err)
	}

	// Persiapan Full Text Search untuk table response
	if err := initializers.DB.Exec("ALTER TABLE responses ADD COLUMN response_vector tsvector").Error; err != nil {
		log.Fatal("err:", err)
	}
	if err := initializers.DB.Exec("UPDATE responses SET response_vector = to_tsvector('indonesian', response)").Error; err != nil {
		log.Fatal("err:", err)
	}
	if err := initializers.DB.Exec("CREATE INDEX idx_response_vector ON responses USING gin(response_vector)").Error; err != nil {
		log.Fatal("err:", err)
	}
}
