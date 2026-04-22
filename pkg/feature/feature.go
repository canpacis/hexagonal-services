package feature

import (
	"github.com/canpacis/hexagonal-services/services/email"
	"github.com/canpacis/hexagonal-services/services/queue"
)

type Service struct {
	QueueService queue.Service
	EmailService email.Service
}

func New(queue queue.Service, email email.Service) *Service {
	return &Service{
		QueueService: queue,
		EmailService: email,
	}
}

func (s *Service) Method() error {
	// Queue a job
	s.QueueService.QueueJob(&queue.QueueJobConfig{})
	// Send email
	s.EmailService.SendEmail(&email.SendEmailConfig{})
	return nil
}
