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

**Create a folder where you want to all you migrations. For e.g: `migrations` and add this file to the folder.**

```go
//file: main.go

package migrations

import "github.com/uptrace/bun/migrate"

// create new migration instance
var Migrations = migrate.NewMigrations()

func init() {
  // register all the migration defined in this folder
	if err := Migrations.DiscoverCaller(); err != nil {
		panic(err)
	}
}
```

**Create a `main.go` in `cmd` folder and add the following code.**

```go
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
go run cmd/main.go migrate
```

## Resources:

- [Uptrace/Bun](https://github.com/uptrace/bun)
- [Urfave/Cli](https://github.com/urfave/cli)
