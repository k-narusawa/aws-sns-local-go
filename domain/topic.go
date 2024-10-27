package domain

type Topic struct {
	TopicArn   string `gorm:"primary_key"`
	Attributes string `gorm:"type:string"`
	Tags       string `gorm:"type:string"`
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
