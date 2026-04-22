package main

import (
	"context"

	"github.com/canpacis/hexagonal-services/pkg/feature"
	"github.com/canpacis/hexagonal-services/services/db"
	"github.com/canpacis/hexagonal-services/services/email"
	"github.com/canpacis/hexagonal-services/services/queue"
)

func main() {
	queries := &db.Queries{
		DB: db.SQLDB{},
	}
	queueService := &queue.SQS{}
	emailService := &email.SES{}

	featureService := feature.New(queries, queueService, emailService)

	ctx := context.Background()
	featureService.Create(ctx)
}
