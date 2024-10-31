package query

import "aws-sns-local-go/usecase/dto"

type TopicQueryService interface {
	FindByTopicArn(topicArn string) ([]dto.TopicDto, error)
}
