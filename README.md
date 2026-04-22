# hexagonal-services

A minimal Go example demonstrating the **Hexagonal Architecture** (Ports and Adapters) pattern.

## Overview

The core idea: business logic depends on *interfaces* (ports), not concrete implementations (adapters). This makes it trivial to swap out infrastructure (e.g. SQS → Kafka) without touching application code.

```
main.go                  — wires adapters into the feature service
pkg/feature/             — business logic; depends only on interfaces
services/queue/          — queue port (interface) + SQS and Kafka adapters
services/email/          — email port (interface) + SES and Mailgun adapters
```

## Usage

```bash
go run main.go
```

To swap implementations, edit `main.go`:

```go
// Use Kafka instead of SQS
queueService := &queue.Kafka{}

// Use Mailgun instead of SES
emailService := &email.Mailgun{}
```

The feature service (`pkg/feature`) requires no changes.

## Key Concept

`feature.Service` receives its dependencies through its constructor:

```go
featureService := feature.New(queueService, emailService)
```

Both `queueService` and `emailService` satisfy interfaces defined in their respective packages. Any struct that implements the interface works — your business logic never knows or cares which adapter is in use.

## Design Principles

### Keep interfaces small

Each port exposes exactly one method. `queue.Service` does one thing: queue a job. `email.Service` does one thing: send an email. This is intentional.

A bloated interface forces every adapter to implement methods it may not need and makes it harder to reason about what a dependency actually requires. A single-method interface is trivial to implement, trivial to mock, and honest about its contract.

```go
// Good — focused, easy to satisfy
type Service interface {
    QueueJob(*QueueJobConfig) error
}

// Avoid — broad interfaces couple adapters to unrelated concerns
type Service interface {
    QueueJob(*QueueJobConfig) error
    DeleteJob(id string) error
    ListJobs() ([]*Job, error)
    PurgeQueue() error
}
```

If a consumer needs multiple capabilities, prefer injecting multiple small interfaces over one large one.

### Define interfaces at the provider level, not the consumer level

The interfaces in this project live in `services/queue` and `services/email` — the same packages that own the config types and adapters. This keeps the contract close to the implementations that must satisfy it.

Defining an interface inside the consumer (`pkg/feature`) would force that package to own types it doesn't control and would scatter the contract across the codebase. When the interface lives with the provider, any package that wants to depend on queue or email behavior has a single, canonical place to look.
