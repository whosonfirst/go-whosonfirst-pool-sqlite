package sqlite

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-pool"
	"github.com/whosonfirst/go-whosonfirst-pool-sqlite/tables"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
	"github.com/whosonfirst/go-whosonfirst-sqlite/database"
	"time"
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

	conn, err := pl.db.Conn()

	if err != nil {
		return -1
	}

	sql := fmt.Sprintf("SELECT COUNT(id) FROM %s", pl.items.Name())

	row := conn.QueryRow(sql)

	var count int64
	err = row.Scan(&count)

	if err != nil {
		return -1
	}

	return count
}

func (pl *SQLiteLIFOPool) Push(i pool.Item) {

	conn, err := pl.db.Conn()

	if err != nil {
		return
	}

	tx, err := conn.Begin()

	if err != nil {
		return
	}

	now := time.Now()
	ts := now.Unix()

	sql := fmt.Sprintf(`INSERT OR REPLACE INTO %s (
		id, created
	) VALUES (
		?, ?
	)`, pl.items.Name())

	_, err = tx.Exec(sql, i.Int(), ts)

	if err != nil {
		return
	}

	tx.Commit()
}

func (pl *SQLiteLIFOPool) Pop() (pool.Item, bool) {

	conn, err := pl.db.Conn()

	if err != nil {
		return nil, false
	}

	tx, err := conn.Begin()

	if err != nil {
		return nil, false
	}

	select_sql := fmt.Sprintf("SELECT MIN(id) FROM %s", pl.items.Name())
	delete_sql := fmt.Sprintf("DELETE FROM %s WHERE id=?", pl.items.Name())

	row := tx.QueryRow(select_sql)

	var int int64

	err = row.Scan(&int)

	if err != nil {
		return nil, false
	}

	_, err = tx.Exec(delete_sql, int)

	if err != nil {
		return nil, false
	}

	err = tx.Commit()

	if err != nil {
		return nil, false
	}

	pi := pool.NewIntItem(int)

	return pi, false
}
