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
	AppID  string `mapstructure:"GROXY_APP_ID"`
	APIKey string `mapstructure:"GROXY_API_Key"`

	Host string `mapstructure:"GROXY_Host"`
	Port int    `mapstructure:"GROXY_Port"`
}

func (c *config) Log() {
	terminalWidth := terminal.LineSize()

	fmt.Println()
	terminal.LineSeparator("-", color.New(color.FgHiCyan), terminalWidth)
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
	fmt.Println()
}

func loadFileConfig() (*config, error) {
	color.HiYellow("Loading configuration from .env file instead of system environment variables due to its validity...")

	var c *config = &config{}

	v := viper.New()
	v.SetEnvPrefix("GROXY")
	v.SetConfigFile(".env")

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			color.HiRed("Couldn't read .env file.")
			return nil, err
		}

		return nil, err
	}

	if err := v.Unmarshal(&c); err != nil {
		color.HiRed(err.Error())
		return nil, err
	}

	return c, nil
}

func loadConfig() (*config, error) {
	color.HiCyan("Starting application...")

	v := viper.New()
	v.SetEnvPrefix("GROXY")
	v.AutomaticEnv()

	// Bind variables
	v.BindEnv("Development")
	v.BindEnv("Host")
	v.BindEnv("Port")

	v.BindEnv("APP_ID")
	v.BindEnv("API_Key")

	valid := true
	for _, key := range v.AllKeys() {
		if v.Get(key) == nil {
			valid = false
		}
	}

	var c *config = &config{
		Development: v.GetBool("Development"),

		AppID:  v.GetString("APP_ID"),
		APIKey: v.GetString("API_Key"),

		Host: v.GetString("Host"),
		Port: v.GetInt("Port"),
	}

	// If the system environment variables can form a valid config file,
	// return it. Otherwise, load from .env file.
	if valid {
		return c, nil
	}

	return loadFileConfig()
}
