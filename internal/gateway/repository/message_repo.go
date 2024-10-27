package repository

import (
	"aws-sns-local-go/domain"

	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db}
}

func (repo *MessageRepository) Save(message domain.Message) error {
	return repo.db.Create(message).Error
}

func (repo *MessageRepository) FindAll() ([]domain.Message, error) {
	var messages []domain.Message
	err := repo.db.Find(&messages).Error
	return messages, err
}
