package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/fabiofcferreira/phstoreapi/terminal"
	"github.com/fatih/color"
)

type config struct {
	Development bool

	// AirTable configuration
	AppID  string
	APIKey string

	PublicHost string
	Host       string
	Port       int
}

func loadConfig(path string) (*config, error) {
	cfg := &config{}

	cfgPath := "config.json"
	if len(path) > 0 {
		cfgPath = path
	}

	f, err := os.Open(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	parser := json.NewDecoder(f)
	err = parser.Decode(&cfg)

	return cfg, nil
}

func (c *config) log() {
	terminalWidth := terminal.TerminalSize()

	color.HiYellow("Configuration file")

	terminal.LineSeparator("-", color.New(color.FgHiCyan), terminalWidth)

	fmt.Printf("Development mode: ")
	terminal.YesNoColored(c.Development)

	fmt.Printf("Host: ")
	color.HiRed(c.Host)

	fmt.Printf("Host: ")
	color.HiRed(strconv.Itoa(c.Port))

	terminal.LineSeparator("-", color.New(color.FgHiCyan), terminalWidth)
}
