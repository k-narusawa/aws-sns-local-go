package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	MessageId              string `gorm:"primary_key"`
	TopicArn               string `gorm:"type:string"`
	TargetArn              string `gorm:"type:string"`
	PhoneNumber            string `gorm:"type:string"`
	Message                string `gorm:"type:string"`
	Subject                string `gorm:"type:string"`
	MessageStructure       string `gorm:"type:string"`
	MessageAttributes      string `gorm:"type:string"`
	MessageDeduplicationId string `gorm:"type:string"`
	MessageGroupId         string `gorm:"type:string"`
	PublishedAt            string `gorm:"type:datetime(6)"`
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
