package aws

type CreateTopicOutput struct {
	CreateTopicResponse CreateTopicResponse
}

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

type ListTopicOutput struct {
	ListTopicResponse ListTopicResponse
}

type ListTopicResponse struct {
	Xmlns            string           `xml:"xmlns,attr"`
	ListTopicsResult ListTopicsResult `xml:"ListTopicsResult"`
	ResponseMetadata struct {
		RequestId string `xml:"RequestId"`
	} `xml:"ResponseMetadata"`
}

type ListTopicsResult struct {
	Topics struct {
		Members []Member `xml:"member"`
	} `xml:"Topics"`
}

type Member struct {
	TopicArn string `xml:"TopicArn"`
}
