package main

import (
	"os"
	"runtime"

	"github.com/fatih/color"
	"github.com/shiena/ansicolor"
	"github.com/sirupsen/logrus"

	"github.com/fabiofcferreira/groxy"
	h "github.com/fabiofcferreira/groxy/http"
)

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

	c, err := loadConfig()
	if err != nil {
		color.HiRed("Couldn't load configuration.")

		log.Fatal(err)
	}

	c.Log()

	cfg := &groxy.Config{
		Development: c.Development,

		AppID:  c.AppID,
		APIKey: c.APIKey,

		Host:       c.Host,
		PublicHost: c.PublicHost,
		Port:       c.Port,
	}

	h.Serve(cfg)
}
