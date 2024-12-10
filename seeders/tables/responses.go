package tables

import (
	"fmt"
	"github.com/arfandyam/Whisper-Me/models/domain"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AddResponses(db *gorm.DB, opt Options) {
	for i := 0; i < opt.Amount; i++ {
		responseId := uuid.New()
		response := domain.Response{
			Id:         responseId,
			QuestionId: opt.QuestionId,
			Response:   faker.Paragraph(),
		}

		if err := db.Create(&response).Error; err != nil {
			fmt.Printf("Error when create question: %s\n", response.Id)
			return
		}
	}
}
