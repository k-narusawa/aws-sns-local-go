package dto

type TopicDto struct {
	TopicArn   string `json:"topic_arn",gorm:"primary_key"`
	Attributes string `json:"attributes",gorm:"type:string"`
	Tags       string `json:"tags",gorm:"type:string"`
}

func (TopicDto) TableName() string {
	return "topics"
}
