package main

import (
	"fmt"
	"os"

	"github.com/docker/go-units"
	"github.com/sirupsen/logrus"
	"github.com/tonicbupt/shinnosuke/cmd/shinnosuke/job"
	"github.com/tonicbupt/shinnosuke/pkg/version"
	"github.com/urfave/cli/v2"
)

func run(c *cli.Context) error {
	var (
		size int64
		err  error
	)
	if c.Args().Len() == 1 {
		size, err = units.FromHumanSize(c.Args().Get(0))
	} else {
		size, err = units.FromHumanSize("100kB")
	}
	if err != nil {
		return err
	}

	logrus.Infof("size: %d", size)
	j := job.NewCompressJob(".", 10, size)
	j.Do()
	logrus.Info("done")
	return nil
}

func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Print(version.Version())
	}

	app := &cli.App{
		Name:    "shinnosuke",
		Usage:   "野原しんのすけ, helps to compress your images in JPEG / PNG",
		Version: version.VERSION,
		Action: func(c *cli.Context) error {
			return run(c)
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatalf("Error compressing: %v", err)
		return
	}
}
