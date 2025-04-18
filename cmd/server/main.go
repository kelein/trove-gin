package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/kelein/trove-gin/cmd/server/wire"
	"github.com/kelein/trove-gin/pkg/config"
	"github.com/kelein/trove-gin/pkg/log"
)

var (
	cfg = flag.String("conf", "config/local.yml", "config file path, eg: -conf ./config/local.yml")
)

// @title           Nunu Example API
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8000
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	flag.Parse()

	conf := config.NewConfig(*cfg)
	logger := log.NewLog(conf)

	// Setup slog default logger
	log.SetupSlog()

	app, cleanup, err := wire.NewWire(conf, logger)
	defer cleanup()
	if err != nil {
		slog.Error("wire injection failed", "error", err)
		os.Exit(1)
	}

	addr := fmt.Sprintf("http://%s:%d", conf.GetString("http.host"), conf.GetInt("http.port"))
	slog.Info("server start listen", "addr", addr)
	slog.Info("swagger docs", "addr", fmt.Sprintf("%s/swagger/index.html", addr))
	if err = app.Run(context.Background()); err != nil {
		slog.Error("server run failed", "error", err)
		os.Exit(1)
	}
}
