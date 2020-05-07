package commands

import (
	"context"
	"path/filepath"
	"strings"
	"time"

	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/photoprism"
	"github.com/photoprism/photoprism/internal/service"
	"github.com/photoprism/photoprism/pkg/fs"
	"github.com/photoprism/photoprism/pkg/txt"
	"github.com/urfave/cli"
)

// PurgeCommand is used to register the index cli command
var PurgeCommand = cli.Command{
	Name:   "purge",
	Usage:  "Removes missing files from search results",
	Flags:  purgeFlags,
	Action: purgeAction,
}

var purgeFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "hard",
		Usage: "delete all data and permanently remove from index",
	},
}

// purgeAction removes missing files from search results
func purgeAction(ctx *cli.Context) error {
	start := time.Now()

	conf := config.NewConfig(ctx)
	service.SetConfig(conf)

	cctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := conf.Init(cctx); err != nil {
		return err
	}

	conf.InitDb()

	// get cli first argument
	subPath := strings.TrimSpace(ctx.Args().First())

	if subPath == "" {
		log.Infof("purging missing files in %s", txt.Quote(filepath.Base(conf.OriginalsPath())))
	} else {
		log.Infof("purging missing files in %s", txt.Quote(fs.RelativeName(filepath.Join(conf.OriginalsPath(), subPath), filepath.Dir(conf.OriginalsPath()))))
	}

	if conf.ReadOnly() {
		log.Infof("read-only mode enabled")
	}

	prg := service.Purge()

	opt := photoprism.PurgeOptions{
		Path: subPath,
		Hard: ctx.Bool("hard"),
	}

	if files, photos, err := prg.Start(opt); err != nil {
		return err
	} else {
		elapsed := time.Since(start)

		log.Infof("purged %d files and %d photos in %s", len(files), len(photos), elapsed)
	}

	conf.Shutdown()

	return nil
}
