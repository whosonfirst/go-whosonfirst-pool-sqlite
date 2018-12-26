package tables

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-pool"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
	"github.com/whosonfirst/go-whosonfirst-sqlite/utils"
	"time"
)

var schema = `CREATE TABLE %s (
	id INTEGER PRIMARY KEY,
        created DATE
);

CREATE INDEX %s_by_date ON %s (created);
`

type IntItemsTable struct {
	ItemsTable
	name string
}

func NewIntItemsTableWithDatabase(db sqlite.Database) (ItemsTable, error) {

	t, err := NewIntItemsTable()

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func NewIntItemsTable() (ItemsTable, error) {

	t := IntItemsTable{
		name: "ints",
	}

	return &t, nil
}

func (t *IntItemsTable) Name() string {
	return t.name
}

func (t *IntItemsTable) Schema() string {
	return fmt.Sprintf(schema, t.Name(), t.Name(), t.Name())
}

func (t *IntItemsTable) InitializeTable(db sqlite.Database) error {
	return utils.CreateTableIfNecessary(db, t)
}

func (t *IntItemsTable) IndexRecord(db sqlite.Database, i interface{}) error {
	return t.Push(db, i.(pool.Item))
}

func (t *IntItemsTable) Push(db sqlite.Database, i pool.Item) error {

	conn, err := db.Conn()

	if err != nil {
		return err
	}

	tx, err := conn.Begin()

	if err != nil {
		return err
	}

	now := time.Now()
	ts := now.Unix()

	sql := fmt.Sprintf(`INSERT OR REPLACE INTO %s (
		id, created
	) VALUES (
		?, ?
	)`, t.Name())

	_, err = tx.Exec(sql, i.Int(), ts)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (t *IntItemsTable) Pop(db sqlite.Database) (pool.Item, bool) {

	conn, err := db.Conn()

	if err != nil {
		return nil, false
	}

	tx, err := conn.Begin()

	if err != nil {
		return nil, false
	}

	select_sql := fmt.Sprintf("SELECT MIN(id) FROM %s", t.Name())
	delete_sql := fmt.Sprintf("DELETE FROM %s WHERE id=?", t.Name())

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
	return pi, true
}

func (t *IntItemsTable) Length(db sqlite.Database) int64 {

	conn, err := db.Conn()

	if err != nil {
		return -1
	}

	sql := fmt.Sprintf("SELECT COUNT(id) FROM %s", t.Name())

	row := conn.QueryRow(sql)

	var count int64
	err = row.Scan(&count)

	if err != nil {
		return -1
	}

	return count
}
