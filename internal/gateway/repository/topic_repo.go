package repository

import (
	"aws-sns-local-go/domain"

	"gorm.io/gorm"
)

type TopicRepository struct {
	db *gorm.DB
}

func NewTopicRepository(db *gorm.DB) *TopicRepository {
	return &TopicRepository{db}
}

func (repo *TopicRepository) Save(topic domain.Topic) error {
	return repo.db.Create(topic).Error
}

func (repo *TopicRepository) FindAll() ([]domain.Topic, error) {
	var topics []domain.Topic
	err := repo.db.Find(&topics).Error
	return topics, err
}
