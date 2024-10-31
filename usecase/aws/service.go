package aws

import (
	"aws-sns-local-go/domain"
	"aws-sns-local-go/usecase/query"

	"github.com/google/uuid"
)

type Service struct {
	TopicRepo   domain.TopicRepository
	MessageRepo domain.MessageRepository
	TopicQSvc   query.TopicQueryService
}

func NewService(
	topicRepo domain.TopicRepository,
	messageRepo domain.MessageRepository,
	topicQSvc query.TopicQueryService,
) *Service {
	return &Service{
		TopicRepo:   topicRepo,
		MessageRepo: messageRepo,
		TopicQSvc:   topicQSvc,
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

func (s *Service) ListTopics() (ListTopicOutput, error) {
	topics, err := s.TopicRepo.FindAll()
	if err != nil {
		return ListTopicOutput{}, err
	}

	topicsArns := []string{}
	for _, topic := range topics {
		topicsArns = append(topicsArns, topic.TopicArn)
	}

	out := ListTopicOutput{}
	resp := ListTopicResponse{}
	resp.Xmlns = "http://sns.amazonaws.com/doc/2010-03-31/"
	resp.ResponseMetadata.RequestId = uuid.New().String()
	for _, topicArn := range topicsArns {
		member := Member{TopicArn: topicArn}
		resp.ListTopicsResult.Topics.Members = append(resp.ListTopicsResult.Topics.Members, member)
	}
	out.ListTopicResponse = resp

	return out, nil
}

func (s *Service) Publish(in PublishInput) (PublishOutput, error) {
	message := domain.NewMessage(
		in.TopicArn,
		in.TargetArn,
		in.PhoneNumber,
		in.Message,
		in.Subject,
		in.MessageStructure,
		in.MessageAttributes,
		in.MessageDeduplicationId,
		in.MessageGroupId,
	)

	if message.PhoneNumber == "" {
		topic, err := s.TopicQSvc.FindByTopicArn(in.TopicArn)
		if err != nil || topic == nil {
			return PublishOutput{}, domain.ErrTopicNotFound
		}
	}

	err := s.MessageRepo.Save(*message)
	if err != nil {
		return PublishOutput{}, err
	}

	out := PublishOutput{}
	resp := PublishResponse{}
	resp.Xmlns = "http://sns.amazonaws.com/doc/2010-03-31/"
	resp.PublishResult = append(resp.PublishResult, PublishResult{MessageId: message.MessageId})
	resp.ResponseMetadata.RequestId = uuid.New().String()
	out.PublishResponse = resp

	return out, nil
}
