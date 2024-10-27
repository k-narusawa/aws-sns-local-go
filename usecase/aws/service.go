package aws

import (
	"aws-sns-local-go/domain"

	"github.com/google/uuid"
)

type Service struct {
	TopicRepo domain.TopicRepository
}

func NewService(topicRepo domain.TopicRepository) *Service {
	return &Service{
		TopicRepo: topicRepo,
	}
}

func (s *Service) CreateTopic(in CreateTopicInput) (CreateTopicResponse, error) {
	topic := domain.NewTopic(
		in.Name,
		in.Attributes,
		in.Tags,
	)

	err := s.TopicRepo.Save(*topic)
	if err != nil {
		return CreateTopicResponse{}, err
	}

	out := CreateTopicResponse{}
	out.Xmlns = "http://sns.amazonaws.com/doc/2010-03-31/"
	out.CreateTopicResult = append(out.CreateTopicResult, CreateTopicResult{TopicArn: topic.TopicArn})
	out.ResponseMetadata.RequestId = uuid.New().String()

	return out, nil
}
