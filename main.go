package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bhumong/dbo-test/config"
	"github.com/bhumong/dbo-test/db"
	"github.com/bhumong/dbo-test/migration"
	"github.com/bhumong/dbo-test/server"
	"github.com/bhumong/dbo-test/util/jwt"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	jwt.Init()
	migration.AutoMigration(db.GetDB())
	server.Init()
}
