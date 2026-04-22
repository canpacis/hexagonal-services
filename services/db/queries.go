package db

import "context"

type Queries struct {
	DB DB // Queries does not know anything about the real implementation
}

func New(db DB) *Queries {
	return &Queries{DB: db}
}

type CreateResourceParams struct{}

func (q *Queries) CreateResource(ctx context.Context, p *CreateResourceParams) error {
	_, err := q.DB.Query(`INSERT INTO feature_resources () VALUES () RETURNING *;`)
	return err
}

func (q *Queries) DeleteResource(ctx context.Context, id string) error {
	_, err := q.DB.ExecContext(ctx, `DELETE FROM feature_resources WHERE id = $1;`, id)
	return err
}
