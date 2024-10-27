package domain

type Topic struct {
	TopicArn   string `json:"topic_arn",gorm:"primary_key"`
	Attributes string `json:"attributes",gorm:"type:string"`
	Tags       string `json:"tags",gorm:"type:string"`
}

func NewTopic(name string, attributes, tags string) *Topic {
	region := "us-west-2"
	accountId := "123456789012"
	arn := "arn:aws:sns:" + region + ":" + accountId + ":" + name
	return &Topic{
		TopicArn:   arn,
		Attributes: attributes,
		Tags:       tags,
	}
}
