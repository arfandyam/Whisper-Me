package repository

import (
	"github.com/arfandyam/Whisper-Me/models/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ResponseRepository struct{}

func NewResponseRepository() ResponseRepositoryInterface {
	return &ResponseRepository{}
}

func (repository *ResponseRepository) CreateResponse(tx *gorm.DB, response *domain.Response) (*domain.Response, error) {
	sql := "INSERT INTO responses (id, question_id, response, created_at, updated_at, deleted_at) VALUES (?, ?, ?, NOW(), ?, ?) returning id, question_id, response"

	rows := tx.Raw(sql, response.Id, response.QuestionId, response.Response, nil, nil).Row()
	if err := rows.Scan(&response.Id, &response.QuestionId, &response.Response); err != nil {
		return nil, err
	}

	return response, nil
}

func (repository *ResponseRepository) FindResponseByQuestionId(tx *gorm.DB, questionId uuid.UUID, fetchPerPage int, cursor *uuid.UUID) ([]domain.Response, error) {
	responses := []domain.Response{}
	var sql string
	var rows *gorm.DB
	if cursor == nil {
		sql = `SELECT id, question_id, response, created_at 
		FROM responses 
		WHERE question_id = ?
		ORDER BY id ASC
		LIMIT ?`
		rows = tx.Raw(sql, questionId, fetchPerPage+1)
	} else {
		sql = `SELECT id, question_id, response, created_at 
		FROM responses 
		WHERE question_id = ? AND id >= ?
		ORDER BY id ASC
		LIMIT ?`
		rows = tx.Raw(sql, questionId, cursor, fetchPerPage+1)
	}

	if err := rows.Scan(&responses).Error; err != nil {
		return nil, err
	}

	return responses, nil
}

func (repository *ResponseRepository) FindPrevCursorResponse(tx *gorm.DB, questionId uuid.UUID, fetchPerPage int, cursor *uuid.UUID) (*uuid.UUID, error) {
	var prevDatas []uuid.UUID

	if cursor == nil {
		return nil, nil
	}

	sql := `SELECT id
	FROM responses
	WHERE question_id = ? AND id < ?
	ORDER BY id DESC
	LIMIT ?`

	rows := tx.Raw(sql, questionId, cursor, fetchPerPage)

	if err := rows.Scan(&prevDatas).Error; err != nil {
		return nil, err
	}

	if len(prevDatas) == 0 {
		return nil, nil
	}

	return &prevDatas[len(prevDatas)-1], nil
}

func (repository *ResponseRepository) SearchResponsesByKeyword(tx *gorm.DB, questionId uuid.UUID, fetchPerPage int, keyword string, rank *float64) ([]domain.Response, error) {
	responses := []domain.Response{}
	var sql string
	var rows *gorm.DB

	if rank == nil {
		sql = `SELECT id, question_id, response, created_at, ts_rank(response_vector, plainto_tsquery(?)) AS rank
		FROM responses 
		WHERE question_id = ? AND response_vector @@ plainto_tsquery(?)
		ORDER BY ts_rank(response_vector, plainto_tsquery(?)) DESC
		LIMIT ?`
		rows = tx.Raw(sql, keyword, questionId, keyword, keyword, fetchPerPage+1)
	} else {
		sql = `SELECT id, question_id, response, created_at, ts_rank(response_vector, plainto_tsquery(?)) AS rank
		FROM responses 
		WHERE question_id = ? AND response_vector @@ plainto_tsquery(?) AND ts_rank(response_vector, plainto_tsquery(?)) <= ?
		ORDER BY ts_rank(response_vector, plainto_tsquery(?)) DESC
		LIMIT ?`
		rows = tx.Raw(sql, keyword, questionId, keyword, keyword, rank, keyword, fetchPerPage+1)
	}

	if err := rows.Scan(&responses).Error; err != nil {
		return nil, err
	}

	return responses, nil
}

func (repository *ResponseRepository) FindPrevRankResponse(tx *gorm.DB, questionId uuid.UUID, fetchPerPage int, keyword string, rank *float64) (*float64, error) {
	var prevDatas []float64

	if rank == nil {
		return nil, nil
	}

	sql := `SELECT ts_rank(response_vector, plainto_tsquery(?)) AS rank
	FROM responses
	WHERE question_id = ? AND ts_rank(response_vector, plainto_tsquery(?)) > ?
	ORDER BY ts_rank(response_vector, plainto_tsquery(?)) ASC
	LIMIT ?`

	rows := tx.Raw(sql, keyword, questionId, keyword, rank, keyword, fetchPerPage)

	if err := rows.Scan(&prevDatas).Error; err != nil {
		return nil, err
	}

	if len(prevDatas) == 0 {
		return nil, nil
	}

	return &prevDatas[len(prevDatas)-1], nil
}
