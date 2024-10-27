package domain

type TopicRepository interface {
	Save(topic Topic) error
	FindAll() ([]Topic, error)
}
