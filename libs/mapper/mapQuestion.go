package mapper

import (
	"github.com/arfandyam/Whisper-Me/models/domain"
	"github.com/arfandyam/Whisper-Me/models/dto"
)

func MapQuestionDomainToQuestionDTO(questions domain.Question) dto.QuestionDTO {
	return dto.QuestionDTO{
		Id: questions.Id,
		UserId: questions.UserId,
		Slug: questions.Slug,
		Topic: questions.Topic,
		Question: questions.Question,
	}
}