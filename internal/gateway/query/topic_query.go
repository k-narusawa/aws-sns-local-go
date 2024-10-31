package query

import (
	"aws-sns-local-go/usecase/dto"

	"gorm.io/gorm"
)

type TopicQueryService struct {
	db *gorm.DB
}

func NewTopicQueryService(db *gorm.DB) *TopicQueryService {
	return &TopicQueryService{db}
}

func (s *TopicQueryService) FindByTopicArn(topicArn string) ([]dto.TopicDto, error) {
	// topicArnの指定がない場合は全て返す
	if topicArn == "" {
		var topics []dto.TopicDto
		err := s.db.Find(&topics).Error
		if err != nil {
			return nil, err
		}
		return topics, nil
	}

	var topic dto.TopicDto
	err := s.db.Where("topic_arn = ?", topicArn).First(&topic).Error
	if err != nil {
		return nil, err
	}
	return []dto.TopicDto{topic}, nil
}
