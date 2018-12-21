package main

import (
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-pool"
	"github.com/whosonfirst/go-whosonfirst-pool-sqlite"
	"log"
)

func main() {

	var dsn = flag.String("dsn", ":memory:", "A valid SQLite DSN")
	flag.Parse()

	p, err := sqlite.NewSQLiteLIFOPool(*dsn)

	if err != nil {
		log.Fatal(err)
	}

	f := pool.NewIntItem(int64(123))

	p.Push(f)
	v, _ := p.Pop()

	fmt.Printf("%d", v.Int())
}
