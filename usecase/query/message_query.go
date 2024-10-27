package query

import "aws-sns-local-go/usecase/dto"

type IMessageQuery interface {
	FindAll(limit, offset int) ([]dto.MessageDto, error)
}
