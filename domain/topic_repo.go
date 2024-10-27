package domain

type TopicRepository interface {
	Save(topic Topic) error
	FindByArn(arn string) (*Topic, error)
	FindAll() ([]Topic, error)
}
