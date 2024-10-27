package aws

import (
	"aws-sns-local-go/domain"

	"github.com/google/uuid"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateTopic(in CreateTopicInput) (CreateTopicResponse, error) {
	topic := domain.NewTopic(
		in.Name,
		in.Attributes,
		in.Tags,
	)

	out := CreateTopicResponse{}
	out.Xmlns = "http://sns.amazonaws.com/doc/2010-03-31/"
	out.CreateTopicResult = append(out.CreateTopicResult, CreateTopicResult{TopicArn: topic.TopicArn})
	out.ResponseMetadata.RequestId = uuid.New().String()

	return out, nil
}
