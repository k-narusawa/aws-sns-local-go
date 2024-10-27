package query

import (
	"aws-sns-local-go/usecase/dto"

	"gorm.io/gorm"
)

type MessageQueryService struct {
	db *gorm.DB
}

func NewMessageQueryService(db *gorm.DB) *MessageQueryService {
	return &MessageQueryService{db}
}

func (s *MessageQueryService) FindAll(limit, offset int) ([]dto.MessageDto, error) {
	var messages []dto.MessageDto
	err := s.db.Limit(limit).Offset(offset).Order("published_at desc").Find(&messages).Error
	return messages, err
}
