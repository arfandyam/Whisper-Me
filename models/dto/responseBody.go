package dto

import "github.com/google/uuid"

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type CreateResponse struct {
	*Response
	Id uuid.UUID `json:"id"`
}

type PageCursorInfo struct {
	NextCursor *uuid.UUID `json:"next_cursor"`
	PrevCursor *uuid.UUID `json:"prev_cursor"`
}

type PageRankInfo struct {
	NextCursor *uuid.UUID `json:"next_cursor"`
	PrevCursor *uuid.UUID `json:"prev_cursor"`
	NextRank   *float64   `json:"next_rank"`
	PrevRank   *float64   `json:"prev_rank"`
}
