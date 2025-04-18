package main

import (
	"context"
	"flag"

	"github.com/kelein/trove-gin/cmd/task/wire"
	"github.com/kelein/trove-gin/pkg/config"
	"github.com/kelein/trove-gin/pkg/log"
)

var (
	cfg = flag.String("conf", "config/local.yml", "config file path, eg: -conf ./config/local.yml")
)

func main() {
	flag.Parse()

	conf := config.NewConfig(*cfg)
	logger := log.NewLog(conf)
	logger.Info("start task")
	app, cleanup, err := wire.NewWire(conf, logger)
	defer cleanup()
	if err != nil {
		panic(err)
	}
	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}

}
