package main

import (
	"flag"
	"os"
	"runtime"

	"github.com/shiena/ansicolor"
	"github.com/sirupsen/logrus"

	"github.com/fabiofcferreira/groxy"
	h "github.com/fabiofcferreira/groxy/http"
)

var cfgPath string

func init() {
	flag.StringVar(&cfgPath, "config", "config.json", "Configuration file")

	flag.Parse()
}

func main() {
	// Execute with all of the CPUs available
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Create logger
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	log.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
	log.SetLevel(logrus.InfoLevel)

	// Load config
	c, err := loadConfig(cfgPath)
	if err != nil {
		log.Error("Could not read the configuration file.")
	}

	// Log configuration values
	c.log()

	cfg := &groxy.Config{
		Development: c.Development,

		AppID:  c.AppID,
		APIKey: c.APIKey,

		PublicHost: c.PublicHost,
		Host:       c.Host,
		Port:       c.Port,
	}

	h.Serve(cfg)
}
