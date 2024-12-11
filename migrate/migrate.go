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
	if err := initializers.DB.Exec("UPDATE questions SET question_vector = to_tsvector('indonesian', COALESCE(topic, '') || ' ' || COALESCE(question, ''))").Error; err != nil {
		log.Fatal("err:", err)
	}
	if err := initializers.DB.Exec("CREATE INDEX idx_question_vector ON questions USING gin(question_vector)").Error; err != nil {
		log.Fatal("err:", err)
	}

	// Persiapan Full Text Search untuk table response
	if err := initializers.DB.Exec("ALTER TABLE responses ADD COLUMN response_vector tsvector").Error; err != nil {
		log.Fatal("err:", err)
	}
	if err := initializers.DB.Exec("UPDATE responses SET response_vector = to_tsvector('indonesian', COALESCE(response, ''))").Error; err != nil {
		log.Fatal("err:", err)
	}
	if err := initializers.DB.Exec("CREATE INDEX idx_response_vector ON responses USING gin(response_vector)").Error; err != nil {
		log.Fatal("err:", err)
	}

	if err := initializers.DB.Exec(`
    CREATE OR REPLACE FUNCTION update_question_vector()
    RETURNS TRIGGER AS $$
    BEGIN
        NEW.question_vector := to_tsvector('indonesian', COALESCE(NEW.topic, '') || ' ' || COALESCE(NEW.question, ''));
        RETURN NEW;
    END;
    $$ LANGUAGE plpgsql;
`).Error; err != nil {
    log.Fatal("Error creating trigger function for questions: ", err)
}

if err := initializers.DB.Exec(`
    CREATE TRIGGER trigger_update_question_vector
    BEFORE INSERT OR UPDATE ON questions
    FOR EACH ROW EXECUTE FUNCTION update_question_vector();
`).Error; err != nil {
    log.Fatal("Error creating trigger for questions: ", err)
}

// For responses
if err := initializers.DB.Exec(`
    CREATE OR REPLACE FUNCTION update_response_vector()
    RETURNS TRIGGER AS $$
    BEGIN
        NEW.response_vector := to_tsvector('indonesian', COALESCE(NEW.response, ''));
        RETURN NEW;
    END;
    $$ LANGUAGE plpgsql;
`).Error; err != nil {
    log.Fatal("Error creating trigger function for responses: ", err)
}

if err := initializers.DB.Exec(`
    CREATE TRIGGER trigger_update_response_vector
    BEFORE INSERT OR UPDATE ON responses
    FOR EACH ROW EXECUTE FUNCTION update_response_vector();
`).Error; err != nil {
    log.Fatal("Error creating trigger for responses: ", err)
}
}
