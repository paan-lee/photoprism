package commands

import (
	"context"
	"time"

	"github.com/urfave/cli"

	"github.com/photoprism/photoprism/internal/config"
)

// MigrateCommand registers the migrate cli command.
var MigrateCommand = cli.Command{
	Name:   "migrate",
	Usage:  "Initializes the index database if needed",
	Action: migrateAction,
}

// migrateAction initializes and migrates the database.
func migrateAction(ctx *cli.Context) error {
	start := time.Now()

	conf := config.NewConfig(ctx)

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := conf.Init(); err != nil {
		return err
	}

	log.Infoln("migrating database")

	conf.InitDb()

	elapsed := time.Since(start)

	log.Infof("database migration completed [%s]", elapsed)

	conf.Shutdown()

	return nil
}
