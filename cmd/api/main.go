package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/matryer/way"

	"github.com/martinrue/parol-api/api"
	"github.com/martinrue/parol-api/logger"
	"github.com/martinrue/parol-api/services"
)

const usage = `Parol API ($SHA)

Usage:
  api --conf=<config-file-path>
`

var conf = flag.String("conf", "", "")

func main() {
	printUsage := func() {
		fmt.Fprint(os.Stderr, strings.Replace(usage, "$SHA", api.Commit, -1))
	}

	flag.Usage = func() {
		printUsage()
		os.Exit(0)
	}

	flag.Parse()

	config, err := readConfig(*conf)
	if err != nil {
		printUsage()
		fmt.Fprintf(os.Stderr, "\nerror → %s\n", err)
		os.Exit(1)
	}

	logger := logger.New("[parol-api]")

	services := &services.Services{
		Config: &services.Config{
			AWSKey:    config.AWSKey,
			AWSSecret: config.AWSSecret,
		},
		Logger: logger,
	}

	server := &api.Server{
		Development: config.Development,
		Logger:      logger,
		Router:      way.NewRouter(),
		Services:    services,
	}

	logger.System("starting server → %s", config.Bind)

	if err := server.Start(config.Bind); err != nil {
		logger.System("server start failed → %s", err)
		os.Exit(3)
	}
}
