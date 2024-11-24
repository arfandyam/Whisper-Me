package repository

import (
	"github.com/arfandyam/Whisper-Me/models/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuestionRepository struct{}

func NewQuestionRepository() QuestionRepositoryInterface {
	return &QuestionRepository{}
}

func (repository *QuestionRepository) CreateQuestion(tx *gorm.DB, userId uuid.UUID, question *domain.Question) (*domain.Question, error) {
	sql := "INSERT INTO questions (id, user_id, slug, topic, question, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?, ?, NOW(), ? ,?) RETURNING id, slug, topic, question"

	rows := tx.Raw(sql, question.Id, userId, question.Slug, question.Topic, question.Question, nil, nil).Row()
	if err := rows.Scan(&question.Id, &question.Slug, &question.Topic, &question.Question); err != nil {
		return nil, err
	}

	return question, nil
}

func (repository *QuestionRepository) EditQuestion(tx *gorm.DB, question *domain.Question) (*domain.Question, error) {
	sql := "UPDATE questions SET slug = ?, topic = ?, question = ?, updated_at = NOW() WHERE id = ? RETURNING id, slug, topic, question"

	rows := tx.Raw(sql, question.Slug, question.Topic, question.Question, question.Id).Row()
	if err := rows.Scan(&question.Id, &question.Slug, &question.Topic, &question.Question); err != nil {
		return nil, err
	}

	return question, nil
}
