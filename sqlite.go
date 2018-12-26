package sqlite

import (
	"github.com/whosonfirst/go-whosonfirst-pool"
	"github.com/whosonfirst/go-whosonfirst-pool-sqlite/tables"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
	"github.com/whosonfirst/go-whosonfirst-sqlite/database"
)

type SQLiteLIFOPool struct {
	pool.LIFOPool
	db    sqlite.Database
	items tables.ItemsTable
}

func NewSQLiteLIFOPool(dsn string) (pool.LIFOPool, error) {

	db, err := database.NewDB(dsn)

	if err != nil {
		return nil, err
	}

	items, err := tables.NewIntItemsTableWithDatabase(db)

	if err != nil {
		return nil, err
	}

	pl := SQLiteLIFOPool{
		db:    db,
		items: items,
	}

	return &pl, nil
}

func (pl *SQLiteLIFOPool) Length() int64 {
	return pl.items.Length(pl.db)
}

func (pl *SQLiteLIFOPool) Push(pi pool.Item) {
	pl.items.Push(pl.db, pi)
}

func (pl *SQLiteLIFOPool) Pop() (pool.Item, bool) {
	return pl.items.Pop(pl.db)
}
