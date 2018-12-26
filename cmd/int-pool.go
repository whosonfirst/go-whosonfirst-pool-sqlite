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

	f1 := pool.NewIntItem(int64(123))
	f2 := pool.NewIntItem(int64(456))

	p.Push(f1)
	p.Push(f2)

	v, _ := p.Pop()

	fmt.Printf("%d", v.Int())
}
