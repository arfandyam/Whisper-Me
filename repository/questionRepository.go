package repository

import (
	"fmt"

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

func (repository *QuestionRepository) FindQuestionsByUserId(tx *gorm.DB, userId uuid.UUID, cursor *uuid.UUID, fetchPerPage int) ([]domain.Question, error) {
	questions := []domain.Question{}
	var sql string
	var rows *gorm.DB
	if cursor == nil {
		sql = `SELECT id, user_id, slug, topic, question 
		FROM questions WHERE user_id = ? 
		ORDER BY id ASC 
		LIMIT ?`
		rows = tx.Raw(sql, userId, fetchPerPage+1)
	} else {
		sql = `SELECT id, user_id, slug, topic, question 
		FROM questions 
		WHERE user_id = ? AND id >= ?
		ORDER BY id ASC
		LIMIT ?`
		rows = tx.Raw(sql, userId, *cursor, fetchPerPage+1)
	}

	if err := rows.Scan(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}

func (repository *QuestionRepository) FindPrevCursorQuestion(tx *gorm.DB, userId uuid.UUID, cursor *uuid.UUID, fetchPerPage int) (*uuid.UUID, error) {
	var prevDatas []uuid.UUID
	if cursor == nil {
		return nil, nil
	}

	sql := `SELECT id
	FROM questions
	WHERE user_id = ? AND id < ?
	ORDER BY id DESC
	LIMIT ?`

	rows := tx.Raw(sql, userId, *cursor, fetchPerPage)
	fmt.Println("rows", rows)
	if err := rows.Scan(&prevDatas).Error; err != nil {
		return nil, err
	}
	if len(prevDatas) == 0 {
		fmt.Println("no data found")
		return nil, nil
	}

	prevCursor := prevDatas[len(prevDatas)-1]

	return &prevCursor, nil
}

func (repository *QuestionRepository) SearchQuestionsByKeyword(tx *gorm.DB, userId uuid.UUID, fetchPerPage int, keyword string, rank *float64) ([]domain.Question, error) {
	var sql string
	var rows *gorm.DB
	questions := []domain.Question{}

	if rank == nil {
		sql = `SELECT id, user_id, slug, topic, question, ts_rank(question_vector, plainto_tsquery(?)) AS rank
		FROM questions 
		WHERE user_id = ? AND question_vector @@ plainto_tsquery(?)
		ORDER BY ts_rank(question_vector, plainto_tsquery(?)) DESC
		LIMIT ?`
		rows = tx.Raw(sql, keyword, userId, keyword, keyword, fetchPerPage+1)
	} else {
		sql = `SELECT id, user_id, slug, topic, question, ts_rank(question_vector, plainto_tsquery(?)) AS rank
		FROM questions 
		WHERE user_id = ? AND question_vector @@ plainto_tsquery(?) AND ts_rank(question_vector, plainto_tsquery(?)) <= ?
		ORDER BY ts_rank(question_vector, plainto_tsquery(?)) DESC
		LIMIT ?`
		rows = tx.Raw(sql, keyword, userId, keyword, keyword, rank, keyword, fetchPerPage+1)
	}

	if err := rows.Scan(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}

func (repository *QuestionRepository) FindPrevRankQuestion(tx *gorm.DB, userId uuid.UUID, fetchPerPage int, keyword string, rank *float64) (*float64, error) {
	var prevDatas []float64

	if rank == nil {
		return nil, nil
	}

	sql := `SELECT ts_rank(question_vector, plainto_tsquery(?)) AS rank
	FROM questions
	WHERE user_id = ? AND ts_rank(question_vector, plainto_tsquery(?)) > ?
	ORDER BY ts_rank(question_vector, plainto_tsquery(?)) ASC
	LIMIT ?`
	
	rows := tx.Raw(sql, keyword, userId, keyword, rank, keyword, fetchPerPage)

	if err := rows.Scan(&prevDatas).Error; err != nil {
		return nil, err
	}

	if len(prevDatas) == 0 {
		return nil, nil
	}

	return &prevDatas[len(prevDatas)-1], nil
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
