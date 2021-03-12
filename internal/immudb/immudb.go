package immudb

import (
	"context"

	immuclient "github.com/codenotary/immudb/pkg/client"
)

// Engineer -
type Engineer struct {
	ID     string
	Salary int
}

type DB struct {
	Ctx    context.Context
	Client immuclient.ImmuClient
}

func NewImmuDB() (DB, error) {
	// TODO:
}

func (db *DB) UpdateSalary(id string, salary int) error {
	// TODO -
}

func (db *DB) GetVerifiedSalary(id string) (int, error) {
	// TODO -
}
