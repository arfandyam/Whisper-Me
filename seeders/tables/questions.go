package tables

import (
	"fmt"

	"github.com/arfandyam/Whisper-Me/libs"
	"github.com/arfandyam/Whisper-Me/models/domain"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)



func AddQuestions(db *gorm.DB, opt Options) {
	for i := 0; i < opt.Amount; i++ {
		questionId := uuid.New()
		topicFaker := faker.Sentence()
		question := domain.Question{
			Id:       questionId,
			UserId:   opt.UserId,
			Slug:     libs.ToSlug(topicFaker, questionId),
			Question: faker.Paragraph(),
			Topic:    topicFaker,
		}

		if err := db.Create(&question).Error; err != nil {
			fmt.Printf("Error when create question: %s\n", question.Topic)
			return
		}
	}
}
