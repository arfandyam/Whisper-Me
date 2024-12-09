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

func (repository *QuestionRepository) CreateQuestion(tx *gorm.DB, question *domain.Question) (*domain.Question, error) {
	sql := "INSERT INTO questions (id, user_id, slug, topic, question, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?, ?, NOW(), ? ,?) RETURNING id, user_id, slug, topic, question"

	rows := tx.Raw(sql, question.Id, question.UserId, question.Slug, question.Topic, question.Question, nil, nil).Row()
	if err := rows.Scan(&question.Id, &question.UserId, &question.Slug, &question.Topic, &question.Question); err != nil {
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

func (repository *QuestionRepository) FindQuestionById(tx *gorm.DB, questionId uuid.UUID) (*domain.Question, error) {
	question := &domain.Question{}
	sql := "SELECT id, user_id, slug, topic, question FROM questions WHERE id = ?"

	if err := tx.Raw(sql, questionId).First(question).Error; err != nil {
		return nil, err
	}

	return question, nil
}

func (repository *QuestionRepository) FindQuestionsByUserId(tx *gorm.DB, userId uuid.UUID, fetchPerPage int, offset int) ([]domain.Question, error) {
	questions := []domain.Question{}
	sql := "SELECT id, user_id, slug, topic, question FROM questions WHERE user_id = ? LIMIT ? OFFSET ?"

	rows := tx.Raw(sql, userId, fetchPerPage, offset)
	if err := rows.Scan(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}

func (repository *QuestionRepository) FindQuestionOwner(tx *gorm.DB, questionId uuid.UUID) (*uuid.UUID, error) {
	var userId uuid.UUID

	sql := "SELECT user_id FROM questions WHERE id = ?"

	rows := tx.Raw(sql, questionId).Row()
	if err := rows.Scan(&userId); err != nil {
		return nil, err
	}

	return &userId, nil
}

func (repository *QuestionRepository) SearchQuestionsByKeyword(tx *gorm.DB, userId uuid.UUID, keyword string, fetchPerPage int, offset int) ([]domain.Question, error) {
	questions := []domain.Question{}

	sql := `
	SELECT id, user_id, slug, topic, question 
	FROM questions 
	WHERE user_id = ? AND question_vector @@ plainto_tsquery(?)
	ORDER BY ts_rank(question_vector, plainto_tsquery(?)) DESC
	LIMIT ? OFFSET ?
	`
	rows := tx.Raw(sql, userId, keyword, keyword, fetchPerPage, offset)
	if err := rows.Scan(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}
