package aws

type CreateTopicResponse struct {
	Xmlns             string              `xml:"xmlns,attr"`
	CreateTopicResult []CreateTopicResult `xml:"CreateTopicResult"`
	ResponseMetadata  struct {
		RequestId string `xml:"RequestId"`
	} `xml:"ResponseMetadata"`
}

type CreateTopicResult struct {
	TopicArn string `xml:"TopicArn"`
}
