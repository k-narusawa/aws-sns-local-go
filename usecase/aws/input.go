package aws

type CreateTopicInput struct {
	Name                 string
	Attributes           string
	Tags                 string
	DataProtectionPolicy string
}

type PublishInput struct {
	TopicArn               string
	TargetArn              string
	PhoneNumber            string
	Message                string
	Subject                string
	MessageStructure       string
	MessageAttributes      string
	MessageDeduplicationId string
	MessageGroupId         string
}
