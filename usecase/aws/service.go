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

func (s *Service) CreateTopic(in CreateTopicInput) (CreateTopicOutput, error) {
	topic := domain.NewTopic(
		in.Name,
		in.Attributes,
		in.Tags,
	)

	err := s.TopicRepo.Save(*topic)
	if err != nil {
		return CreateTopicOutput{}, err
	}

	out := CreateTopicOutput{}
	resp := CreateTopicResponse{}
	resp.Xmlns = "http://sns.amazonaws.com/doc/2010-03-31/"
	resp.CreateTopicResult = append(resp.CreateTopicResult, CreateTopicResult{TopicArn: topic.TopicArn})
	resp.ResponseMetadata.RequestId = uuid.New().String()
	out.CreateTopicResponse = resp

	return out, nil
}
