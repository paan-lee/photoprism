package commands

import (
	"context"
	"time"

	"github.com/dustin/go-humanize/english"
	"github.com/urfave/cli"

	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/photoprism"
	"github.com/photoprism/photoprism/internal/service"
)

// CleanUpCommand registers the cleanup command.
var CleanUpCommand = cli.Command{
	Name:   "cleanup",
	Usage:  "Removes orphaned index entries, sidecar and thumbnail files",
	Flags:  cleanUpFlags,
	Action: cleanUpAction,
}

var cleanUpFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "dry",
		Usage: "dry run, don't actually remove anything",
	},
}

// cleanUpAction removes orphaned index entries, sidecar and thumbnail files.
func cleanUpAction(ctx *cli.Context) error {
	cleanupStart := time.Now()

	conf := config.NewConfig(ctx)
	service.SetConfig(conf)

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := conf.Init(); err != nil {
		return err
	}

	conf.InitDb()

	if conf.ReadOnly() {
		log.Infof("config: read-only mode enabled")
	}

	w := service.CleanUp()

	opt := photoprism.CleanUpOptions{
		Dry: ctx.Bool("dry"),
	}

	// Start cleanup worker.
	if thumbnails, _, sidecars, err := w.Start(opt); err != nil {
		return err
	} else if total := thumbnails + sidecars; total > 0 {
		log.Infof("removed %s in %s", english.Plural(total, "file", "files"), time.Since(cleanupStart))
	}

	conf.Shutdown()

	return nil
}
