package main

import (
	"fmt"
	"strconv"

	"github.com/fabiofcferreira/groxy/terminal"
	"github.com/fatih/color"
	"github.com/spf13/viper"
)

type config struct {
	Development bool `mapstructure:"GROXY_Development"`

	// AirTable configuration
	AppID  string `mapstructure:"GROXY_App_ID"`
	APIKey string `mapstructure:"GROXY_API_Key"`

	Host       string `mapstructure:"GROXY_Host"`
	PublicHost string `mapstructure:"GROXY_PublicHost"`
	Port       int    `mapstructure:"GROXY_Port"`
}

func (c *config) Log() {
	terminalWidth := terminal.LineSize()

	color.HiYellow("Configuration")
	terminal.LineSeparator("-", color.New(color.FgHiCyan), terminalWidth)

	fmt.Printf("Development mode: ")
	terminal.YesNoColored(c.Development)

	fmt.Printf("Host: ")
	color.HiRed(c.Host)

	fmt.Printf("Host: ")
	color.HiRed(strconv.Itoa(c.Port))

	fmt.Printf("App ID: ")
	color.HiBlue(c.AppID)

	if c.Development {
		fmt.Printf("API key: ")
		color.HiBlue(c.APIKey)
	}

	terminal.LineSeparator("-", color.New(color.FgHiCyan), terminalWidth)
}

func loadConfig() (*config, error) {
	var c *config = &config{}
	var err error

	v := viper.New()

	// Config files
	v.SetConfigFile(".env")
	v.SetEnvPrefix("GROXY")

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	if err = v.Unmarshal(&c); err != nil {
		return nil, err
	}

	return c, err
}
