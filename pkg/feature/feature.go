package feature

import (
	"context"

	"github.com/canpacis/hexagonal-services/services/db"
	"github.com/canpacis/hexagonal-services/services/email"
	"github.com/canpacis/hexagonal-services/services/queue"
)

type Service struct {
	Queries      *db.Queries
	QueueService queue.Service
	EmailService email.Service
}

func New(queries *db.Queries, queue queue.Service, email email.Service) *Service {
	return &Service{
		Queries:      queries,
		QueueService: queue,
		EmailService: email,
	}
}

func (s *Service) Create(ctx context.Context) error {
	if err := s.Queries.CreateResource(ctx, &db.CreateResourceParams{}); err != nil {
		return err
	}
	// Queue a job
	if err := s.QueueService.QueueJob(&queue.QueueJobConfig{}); err != nil {
		return err
	}
	// Send email
	if err := s.EmailService.SendEmail(&email.SendEmailConfig{}); err != nil {
		return err
	}
	return nil
}
