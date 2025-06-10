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

		sql := "INSERT INTO responses (id, question_id, response, created_at, updated_at, deleted_at) VALUES (?, ?, ?, NOW(), ?, ?) RETURNING id"
		rows := db.Raw(sql, response.Id, response.QuestionId, response.Response, nil, nil).Row()
		if err := rows.Scan(&response.Id); err != nil {
			fmt.Printf("Error when create response: %s\n", response.Id)
			return
		}
	}
}
