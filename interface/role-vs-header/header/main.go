package header

import (
	"context"
	"database/sql"
	"database/sql/driver"
)

func Example(db *sql.DB) {
	// The consumer expect an abstract resource theory,
	// but in practice, only *sql.DB can be used as supplier
	//
	// Introducing different implementation requires changes across the system,
	// unless the supplier fake all the really needed functionality
	if err := (Consumer{resource: db}.A()); err != nil {
		panic(err)
	}
}

//----------------------------------------------------------------------------------------------------------------------

// sqlDB is the header interface to *sql.DB
type sqlDB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Begin() (*sql.Tx, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Ping() error
	Driver() driver.Driver
	Close() error
	Conn(ctx context.Context) (*sql.Conn, error)
	PingContext(ctx context.Context) error
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	// ... and so on
}

type Consumer struct {
	resource sqlDB
}

func (m Consumer) A() error {
	var count int
	if err := m.resource.QueryRow(`SELECT 1`).Scan(&count); err != nil {
		return err
	}

	return nil
}
