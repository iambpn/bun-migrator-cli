# bun-migrator-cli

Golang's Bun ORM migration cli initializer.

This is the small wrapper to the Bun migration package. This package includes migration cli configuration with basic migration functionality.

You can read more about the Migration Tool in [Bun Migration Docs](https://bun.uptrace.dev/guide/migrations.html).

Also, Most of the code used here is taken from [Bun Example Project](https://github.com/uptrace/bun/tree/master/example/migrate)

## Usage

```go
package main

import (
  "database/sql"
  "log"
  "os"

  "github.com/uptrace/bun"
  "github.com/uptrace/bun/dialect/sqlitedialect"
  "github.com/uptrace/bun/driver/sqliteshim"
  "github.com/uptrace/bun/migrate"
  migratorCli "github.com/uptrace/bun-migrator-cli"
)

// instantiate DB
sqlDb, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")

if err != nil {
  log.Fatal(err)
}

// get bun db instance
db := bun.NewDB(sqlDb, sqlitedialect.New())

// instantiate Migrator
migrator := migrate.NewMigrator(db, Migrations)

// initialize cli
migratorCli.InitCli(migrator, os.Args)
```

## Resources:

- [Uptrace/Bun](https://github.com/uptrace/bun)
- [Urfave/Cli](https://github.com/urfave/cli)
