package role

import (
	"database/sql"
)

func Example(db *sql.DB, tx *sql.Tx) {
	// Consumer example using *sql.DB supplier
	if err := (Consumer{resource: db}.A()); err != nil {
		panic(err)
	}

	// The same Consumer, but using a different implementation,
	// and working on a transaction, because the supplier can be a *sql.Tx as well
	if err := (Consumer{resource: tx}.A()); err != nil {
		panic(err)
	}
}

//----------------------------------------------------------------------------------------------------------------------

type Resource interface {
	QueryRow(query string, args ...interface{}) *sql.Row
}

type Consumer struct {
	resource Resource
}

func (m Consumer) A() error {
	var count int
	if err := m.resource.QueryRow(`SELECT 1`).Scan(&count); err != nil {
		return err
	}

	return nil
}
