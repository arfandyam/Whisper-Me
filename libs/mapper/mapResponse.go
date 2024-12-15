package mapper

import (
	"github.com/arfandyam/Whisper-Me/models/domain"
	"github.com/arfandyam/Whisper-Me/models/dto"
)

func MapResponseDomainToFullResponseDTO(response domain.Response) dto.FullResponseDTO {
	return dto.FullResponseDTO{
		ResponseDTO: MapResponseDomainToResponseDTO(response),
		ResponseTimestampDTO: MapResponseDomainToResponseTimestampDTO(response),
	}
}

func MapResponseDomainToResponseDTO(response domain.Response) dto.ResponseDTO {
	return dto.ResponseDTO{
		Id:         response.Id,
		QuestionId: response.QuestionId,
		Response:   response.Response,
	}
}

func MapResponseDomainToResponseTimestampDTO(response domain.Response) dto.ResponseTimestampDTO {
	return dto.ResponseTimestampDTO{
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	}
}
