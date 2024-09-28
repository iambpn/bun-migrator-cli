package migratorCli

import (
	"database/sql"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"

	"testing"
)

var Migrations = migrate.NewMigrations()

func init() {
	if err := Migrations.DiscoverCaller(); err != nil {
		panic(err)
	}
}

func getMigrator() *migrate.Migrator {
	sqlDb, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")

	if err != nil {
		log.Fatal(err)
	}

	db := bun.NewDB(sqlDb, sqlitedialect.New())

	return migrate.NewMigrator(db, Migrations)
}

func TestInit(t *testing.T) {
	InitCli(getMigrator(), []string{})
}

func TestGetCommands(t *testing.T) {
	commands := getCommands(getMigrator())

	if len(commands) != 10 {
		t.Errorf("Expected 6 commands, got %d", len(commands))
	}

	commands = getCommands(getMigrator(), &cli.Command{
		Name:  "test",
		Usage: "test",
		Action: func(c *cli.Context) error {
			return nil
		},
	})

	if len(commands) != 11 {
		t.Errorf("Expected 7 commands, got %d", len(commands))
	}
}
