package db

import (
	"context"
	"database/sql"
)

type DB interface {
	Begin() (*sql.Tx, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

// Generic SQL db implementation of the DB interface
type SQLDB struct {
	*sql.DB
}

// Aurora implementation of the DB interface
type AuroraDB struct {
	// Aurora client here
}

func (*AuroraDB) Begin() (*sql.Tx, error) {
	// Use aurora client here
	return nil, nil
}

func (*AuroraDB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	// Use aurora client here
	return nil, nil
}

func (*AuroraDB) Exec(query string, args ...any) (sql.Result, error) {
	// Use aurora client here
	return nil, nil
}

func (*AuroraDB) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	// Use aurora client here
	return nil, nil
}

func (*AuroraDB) Query(query string, args ...any) (*sql.Rows, error) {
	// Use aurora client here
	return nil, nil
}

func (*AuroraDB) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	// Use aurora client here
	return nil, nil
}
