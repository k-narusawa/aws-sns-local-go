package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	MessageId              string `json:"message_id",gorm:"primary_key"`
	TopicArn               string `json:"topic_arn",gorm:"type:string"`
	TargetArn              string `json:"target_arn",gorm:"type:string"`
	PhoneNumber            string `json:"phone_number",gorm:"type:string"`
	Message                string `json:"message",gorm:"type:string"`
	Subject                string `json:"subject",gorm:"type:string"`
	MessageStructure       string `json:"message_structure",gorm:"type:string"`
	MessageAttributes      string `json:"message_attributes",gorm:"type:string"`
	MessageDeduplicationId string `json:"message_deduplication_id",gorm:"type:string"`
	MessageGroupId         string `json:"message_group_id",gorm:"type:string"`
	PublishedAt            string `json:"published_at",gorm:"type:datetime(6)"`
}

func NewMessage(
	topicArn string,
	targetArn string,
	phoneNumber string,
	message string,
	subject string,
	messageStructure string,
	messageAttributes string,
	messageDeduplicationId string,
	messageGroupId string,
) *Message {
	messageId := uuid.New().String()
	return &Message{
		MessageId:              messageId,
		TopicArn:               topicArn,
		TargetArn:              targetArn,
		PhoneNumber:            phoneNumber,
		Message:                message,
		Subject:                subject,
		MessageStructure:       messageStructure,
		MessageAttributes:      messageAttributes,
		MessageDeduplicationId: messageDeduplicationId,
		MessageGroupId:         messageGroupId,
		PublishedAt:            time.Now().Format(time.RFC3339Nano),
	}
}
