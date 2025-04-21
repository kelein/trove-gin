package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/kelein/trove-gin/cmd/server/wire"
	"github.com/kelein/trove-gin/docs"
	"github.com/kelein/trove-gin/pkg/config"
	"github.com/kelein/trove-gin/pkg/log"
)

var (
	cfg = flag.String("conf", "config/dev.yaml", "config file path")
)

func init() { initSwaggerInfo() }

// initSwaggerInfo setup swagger info
//
// * Basic Info
// @license.name Apache 2.0
// @contact.name trove-gin
// @contact.url https://github.com/kelein/trove-gin
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
//
// * Authentication Info
// @securityDefinitions.apiKey Bearer
// @name Authorization
// @in header
func initSwaggerInfo() {
	docs.SwaggerInfo.BasePath = "v1"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = "localhost:8000"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Title = "Trove Gin API Server"
	docs.SwaggerInfo.Description = "Trove Gin API Server"
	docs.SwaggerInfo.InfoInstanceName = "swagger"
}

func main() {
	flag.Parse()

	conf := config.NewConfig(*cfg)
	logger := log.NewLog(conf)

	// Setup slog default logger
	log.SetupSlog(conf)

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
