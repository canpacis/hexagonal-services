package main

import (
	"github.com/canpacis/hexagonal-services/pkg/feature"
	"github.com/canpacis/hexagonal-services/services/email"
	"github.com/canpacis/hexagonal-services/services/queue"
)

func main() {
	queueService := &queue.SQS{}
	emailService := &email.SES{}

	featureService := feature.New(queueService, emailService)

	featureService.Method()
}
