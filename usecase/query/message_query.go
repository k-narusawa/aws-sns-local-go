package query

import "aws-sns-local-go/usecase/dto"

type IMessageQuery interface {
	FindAll(phoneNumber string, limit, offset int) ([]dto.MessageDto, error)
}
