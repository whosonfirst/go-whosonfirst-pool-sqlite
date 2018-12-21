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
	items sqlite.Table
}

func NewSQLiteLIFOPool(dsn string) (pool.LIFOPool, error) {

	db, err := database.NewDB(dsn)

	if err != nil {
		return nil, err
	}

	items, err := tables.NewItemsTableWithDatabase(db)

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

	// please write me
	return -1
}

func (pl *SQLiteLIFOPool) Push(i pool.Item) {

	// error handling/checking?
	pl.items.IndexRecord(pl.db, i)
}

func (pl *SQLiteLIFOPool) Pop() (pool.Item, bool) {

	// please write me
	return nil, false
}
