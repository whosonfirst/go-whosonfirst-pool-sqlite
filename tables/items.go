package tables

import (
	"errors"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-pool"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
	"github.com/whosonfirst/go-whosonfirst-sqlite/utils"
)

var schema = `CREATE TABLE %s (
	id TEXT PRIMARY KEY,
	item TEXT,
        created DATE
);

CREATE INDEX %s_by_date ON %s (created);
`

type ItemsTable struct {
	sqlite.Table
	name string
}

func NewItemsTableWithDatabase(db sqlite.Database) (sqlite.Table, error) {

	t, err := NewItemsTable()

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func NewItemsTable() (sqlite.Table, error) {

	t := ItemsTable{
		name: "items",
	}

	return &t, nil
}

func (t *ItemsTable) Name() string {
	return t.name
}

func (t *ItemsTable) Schema() string {

	// this is a bit stupid really... (20170901/thisisaaronland)
	return fmt.Sprintf(schema, t.Name(), t.Name(), t.Name())
}

func (t *ItemsTable) InitializeTable(db sqlite.Database) error {

	return utils.CreateTableIfNecessary(db, t)
}

func (t *ItemsTable) IndexRecord(db sqlite.Database, i interface{}) error {
	return t.indexItem(db, i.(pool.Item))
}

func (t *ItemsTable) indexItem(db sqlite.Database, i pool.Item) error {

	return errors.New("Please write me")

	conn, err := db.Conn()

	if err != nil {
		return err
	}

	tx, err := conn.Begin()

	if err != nil {
		return err
	}

	// do stuff here...

	return tx.Commit()
}
