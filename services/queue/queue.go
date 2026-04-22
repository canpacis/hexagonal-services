package queue

type Service interface {
	QueueJob(*QueueJobConfig) error
}

type QueueJobConfig struct {
	// Configuration
}

// SQS implementation of the queue service
type SQS struct{}

func (*SQS) QueueJob(*QueueJobConfig) error {
	return nil
}

// Kafka implementation of the queue service
type Kafka struct{}

func (*Kafka) QueueJob(*QueueJobConfig) error {
	return nil
}
