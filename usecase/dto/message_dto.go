package dto

type MessageDto struct {
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

func (MessageDto) TableName() string {
	return "messages"
}
