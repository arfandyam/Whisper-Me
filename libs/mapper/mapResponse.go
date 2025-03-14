package mapper

import (
	"github.com/arfandyam/Whisper-Me/models/domain"
	"github.com/arfandyam/Whisper-Me/models/dto"
)

func MapResponseDomainToResponseDTO(response domain.Response) dto.ResponseDTO {
	return dto.ResponseDTO{
		Id: response.Id,
		QuestionId: response.QuestionId,
		Response: response.Response,
		CreatedAt: response.CreatedAt,
	}
}