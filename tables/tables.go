package tables

import (
	"github.com/whosonfirst/go-whosonfirst-pool"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
)

type ItemsTable interface {
	sqlite.Table
	SQLiteLIFOPoolTable
}

type SQLiteLIFOPoolTable interface {
	Length(sqlite.Database) int64
	Push(sqlite.Database, pool.Item) error
	Pop(sqlite.Database) (pool.Item, bool)
}
