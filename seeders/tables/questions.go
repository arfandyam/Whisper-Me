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
		questionSlug := libs.ToSlug(topicFaker, questionId)
		question := domain.Question{
			Id:       questionId,
			UserId:   opt.UserId,
			Slug:     questionSlug,
			Question: faker.Paragraph(),
			Topic:    topicFaker,
			UrlKey:   libs.SlugToBase62(questionSlug),
		}

		sql := "INSERT INTO questions (id, user_id, slug, topic, question, url_key, created_at, updated_at, deleted_at) VALUES(?, ?, ?, ?, ?, ?, NOW(), ?, ?) RETURNING id"
		rows := db.Raw(sql, question.Id, question.UserId, question.Slug, question.Topic, question.Question, question.UrlKey, nil, nil).Row()
		if err := rows.Scan(&question.Id); err != nil {
			fmt.Printf("Error when create question: %s\n", question.Topic)
			fmt.Println("err:", err)
			return
		}
	}
}
