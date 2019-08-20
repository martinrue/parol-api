package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

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
			AWSRegion: config.AWSRegion,
			AWSBucket: config.AWSBucket,
			Keys:      config.Keys,
		},
		Logger: logger,
		Usage: &services.Usage{
			Started:       time.Now(),
			TotalRequests: 0,
			HourlyUsage:   make(map[string]int, 0),
		},
	}

	server := &api.Server{
		Development: config.Development,
		Router:      way.NewRouter(),
		Services:    services,
	}

	logger.System("starting server → %s", config.Bind)

	if err := server.Start(config.Bind); err != nil {
		logger.System("server start failed → %s", err)
		os.Exit(3)
	}
}
