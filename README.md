# bun-migrator-cli

Golang's Bun ORM migration cli initializer.

This is the small wrapper to the Bun migration package. This package includes migration cli configuration with basic migration functionality.

You can read more about the Migration Tool in [Bun Migration Docs](https://bun.uptrace.dev/guide/migrations.html).

Also, Most of the code used here is taken from [Bun Example Project](https://github.com/uptrace/bun/tree/master/example/migrate)

### Install

```bash
go get github.com/iambpn/bun-migrator-cli
```

## Usage

**Create a folder where you want to add your migration files. For e.g: In `/migrations` folder, add this file.**

```go
//file: migrations/init.go

package migrations

import "github.com/uptrace/bun/migrate"

// create new migration instance
var Migrations = migrate.NewMigrations()

// Run when the module is initialized
func init() {
  // register all the sql migration defined in this folder
	if err := Migrations.DiscoverCaller(); err != nil {
		panic(err)
	}
}
```

**Create a `main.go` in `cmd` folder and add this code to setup migrator cli.**

```go
// file: cmd/main.go
package main

import (
  "database/sql"
  "log"
  "os"

  ".../migrations"
  "github.com/uptrace/bun"
  "github.com/uptrace/bun/dialect/sqlitedialect"
  "github.com/uptrace/bun/driver/sqliteshim"
  "github.com/uptrace/bun/migrate"
  migratorCli "github.com/iambpn/bun-migrator-cli"
)

func main(){
  // instantiate raw DB
  sqlDb, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")

  if err != nil {
    log.Fatal(err)
  }

  // get bun db instance
  db := bun.NewDB(sqlDb, sqlitedialect.New())

  // instantiate Migrator
  migrator := migrate.NewMigrator(db, migrations.Migrations)

  // initialize cli
  migratorCli.InitCli(migrator, os.Args)
}
```

**Run migrator cli**

```bash
go run cmd/main.go help # to see the help message
go run cmd/main.go init # to initialize the migration table in database
go run cmd/main.go create <migration_name> # to create a new migration file
go run cmd/main.go migrate # to run all the pending migrations
```

## Resources:

- [Uptrace/Bun](https://github.com/uptrace/bun)
- [Urfave/Cli](https://github.com/urfave/cli)
